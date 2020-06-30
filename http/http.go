package http

import (
	"errors"
	"io/ioutil"
	"strconv"

	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/kexirong/msg-sender/email"
	"github.com/kexirong/msg-sender/wechat"
)

var (
	//Info log level
	Info *log.Logger
	//Warning log level
	Warning *log.Logger
	//Error log level
	Error *log.Logger
)

type logfile struct {
	file        *os.File
	isOpen      bool
	rolte       <-chan time.Time
	preFileName string
}

func (lf *logfile) Write(b []byte) (int, error) {
	select {
	case <-lf.rolte:

		err := lf.SetFile(lf.preFileName + time.Now().Format("2006-01-02") + ".log")
		if err != nil {
			return 0, err
		}

	default:
		if !lf.isOpen {
			return 0, errors.New("logfile is not config")
		}
	}

	return lf.file.Write(b)
}

func newLogFile(preFileName string) *logfile {
	var lf logfile
	lf.preFileName = preFileName
	err := lf.SetFile(lf.preFileName + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		panic(err)
	}

	go func() {
		t := time.Now()
		nx := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
		<-time.After(nx.Sub(t))
		lf.SetFile(lf.preFileName + time.Now().Format("2006-01-02") + ".log")
		lf.rolte = time.Tick(time.Hour * 24)
	}()

	return &lf
}

func (lf *logfile) SetFile(filename string) error {
	var filepath string
	wd, err := os.Getwd()
	if err == nil {
		filepath = path.Join(wd, "log/"+filename)
	} else {
		panic(err)
	}

	if pathIsExist(filepath) {
		err := os.Rename(filepath, filepath+strconv.FormatInt(time.Now().Unix(), 10))
		if err != nil {
			return err
		}
	}
	if lf.isOpen {
		lf.Close()
	}
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		return err
	}
	lf.file = file
	lf.isOpen = true
	return nil
}

func (lf *logfile) Close() {
	lf.file.Close()
}

func init() {
	if !pathIsExist("./log") {
		err := os.Mkdir("./log", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	fli := newLogFile("msg-sender")
	Info = log.New(fli, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(fli, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(fli, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func pathIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//SrvStart start http lisent server
func SrvStart(cfg *simplejson.Json) {
	//fmt.Println(cfg)

	httpAddr := cfg.Get("http").Get("listen").MustString("")
	smtpCfg := cfg.Get("smtp")
	smtpAddr := smtpCfg.Get("address").MustString("")
	smtpUser := smtpCfg.Get("username").MustString("")
	smtpPass := smtpCfg.Get("password").MustString("")
	smtpAuthtype := smtpCfg.Get("authtype").MustString("")
	wxCfg := cfg.Get("wechat")
	wxCorpID := wxCfg.Get("CorpID").MustString("")
	wxAgentID := wxCfg.Get("AgentId").MustInt(0)
	wxSecret := wxCfg.Get("Secret").MustString("")

	Info.Println(fmt.Sprintf("httpAddr:%s", httpAddr))
	Info.Println(fmt.Sprintf("smtpAddr:%s,smtpUser:%s,smtpPass:%s", smtpAddr, smtpUser, smtpPass))

	s := email.New(smtpAddr, smtpUser, smtpPass, smtpAuthtype)
	wx := wechat.New(wxCorpID, wxAgentID, wxSecret)
	err := wx.GetAccToken()
	if err != nil {
		Error.Printf("getAccToken failed: %s/n", err)
	}
	Info.Printf("getAccToken done: %s /n", wx.AccToken)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		/* fmt.Println(r.URL.Query())
		   fmt.Println(r.Method)
		   err := r.ParseForm()
		   if err != nil {
		       panic(err)
		   }
		   fmt.Println(r.Form.Get("name"))
		   fmt.Println(r.Form["name"])*/
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(testStr))
	})

	http.HandleFunc("/sender/mail", func(w http.ResponseWriter, r *http.Request) {
		var pload payload
		contentType := r.Header.Get("Content-Type")
		switch {
		case strings.Contains(contentType, "json"):
			w.Header().Set("Content-Type", contentType)
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				Error.Println(err)
			}
			err = pload.UnmarshalJSON(b)
			if err != nil {
				Error.Println(err)
			}

		case strings.Contains(contentType, "x-www-form-urlencoded"):
			err := r.ParseForm()
			if err != nil {
				Error.Println(err)
			}
			pload.From = r.PostFormValue("from")
			pload.To = r.PostFormValue("to")
			pload.Subject = r.PostFormValue("subject")
			pload.Content = r.PostFormValue("content")
			pload.ContentType = r.PostFormValue("content_type")
		default:
			Error.Println("invalid Content-Type")
		}
		Info.Printf("#sendMail# client:%s, %s-->%s, requestType:%s, subject:%s, content:%s\n", r.RemoteAddr, pload.From, pload.To, contentType, pload.Subject, pload.Content)

		err = s.SendMail(pload.From, strings.Split(pload.To, ","), pload.Subject, pload.Content, pload.ContentType)

		if err != nil {
			Error.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write([]byte(`{"errcode": 0,"errmsg":"ok"}`))
		}
	})

	http.HandleFunc("/sender/wechat", func(w http.ResponseWriter, r *http.Request) {
		var pload payload
		contentType := r.Header.Get("Content-Type")
		switch {
		case strings.Contains(contentType, "json"):
			w.Header().Set("Content-Type", contentType)
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				Error.Println(err)
			}
			err = pload.UnmarshalJSON(b)
			if err != nil {
				Error.Println(err)
			}

		case strings.Contains(contentType, "x-www-form-urlencoded"):
			err := r.ParseForm()
			if err != nil {
				Error.Println(err)
			}
			pload.To = r.PostFormValue("to")
			pload.Content = r.PostFormValue("content")
		default:
			Error.Println("invalid Content-Type")
		}
		Info.Printf("#sendWechat# client:%s, to:%s, requestType:%s, content:%s\n", r.RemoteAddr, pload.To, contentType, pload.Content)

		resp, err := wx.SendMsg(pload.To, "", pload.Content)
		if err != nil {
			Error.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			Info.Printf(string(resp))
			w.Write(resp)
		}
	})

	log.Fatal(http.ListenAndServe(httpAddr, nil))

}
