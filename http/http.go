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

	"anymsg/dingding"
	"anymsg/email"
	"anymsg/wechat"
	"github.com/bitly/go-simplejson"
)

var (
	//Info log level
	Info *log.Logger
	//Warning log level
	Warning *log.Logger
	//Error log level
	Error   *log.Logger
	MsgType string
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
	dingCfg := cfg.Get("dingding")
	dingurl := dingCfg.Get("Url").MustString("")
	dingToken := dingCfg.Get("AccessToken").MustString("")

	Info.Println(fmt.Sprintf("httpAddr:%s", httpAddr))
	Info.Println(fmt.Sprintf("smtpAddr:%s,smtpUser:%s,smtpPass:%s", smtpAddr, smtpUser, smtpPass))

	s := email.New(smtpAddr, smtpUser, smtpPass, smtpAuthtype)
	wx := wechat.New(wxCorpID, wxAgentID, wxSecret)
	err := wx.GetAccToken()
	if err != nil {
		Error.Printf("getAccToken failed: %s/n", err)
	}
	dingding := dingding.New(dingurl, dingToken)
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
		fmt.Println("contentType: ", contentType)
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
			fmt.Println("pload: ", pload)

		case strings.Contains(contentType, "x-www-form-urlencoded"):
			err := r.ParseForm()
			if err != nil {
				Error.Println(err)
			}
			pload.To = r.PostFormValue("to")

			//testmarkdown := r.PostFormValue("markdown")
			//if MsgType == "markdown"{
			//	var markdown Markdown
			//	err = json.Unmarshal([]byte(testmarkdown), &markdown)
			//	if err != nil{
			//		fmt.Println("json.Unmarshal err: ",err)
			//	}
			//	pload.Content = r.PostFormValue("content")
			//	pload.Markdown = markdown
			//}else {
			fmt.Println("test: ", MsgType)
			pload.Content = r.PostFormValue("content")
			fmt.Println("pload.Content: ", pload.Content)
			//}

		case strings.Contains(contentType, "markdown"):
			fmt.Println("r.Body: ", contentType)
			w.Header().Set("Content-Type", contentType)
			b, err := ioutil.ReadAll(r.Body)
			fmt.Println("r.Body: ", b)
			if err != nil {
				Error.Println(err)
			}
			err = pload.UnmarshalJSON(b)
			if err != nil {
				Error.Println(err)
			}

		default:
			fmt.Println("err")
			Error.Println("invalid Content-Type")
		}
		Info.Printf("#sendWechat# client:%s, to:%s, requestType:%s, content:%s\n", r.RemoteAddr, pload.To, contentType, pload.Content)
		// 判断发送内容是否是markdown格式的，默认非markdown格式
		if pload.MsgType == "markdown" {
			resp, err := wx.SendMsgMarkdown(pload.To, "", pload.Content)
			if err != nil {
				Error.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				Info.Printf(string(resp))
				w.Write(resp)
			}
		} else {
			resp, err := wx.SendMsg(pload.To, "", pload.Content)
			if err != nil {
				Error.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				Info.Printf(string(resp))
				w.Write(resp)
			}
		}

	})
	http.HandleFunc("/sender/dingding", func(w http.ResponseWriter, r *http.Request) {
		var pload payload
		contentType := r.Header.Get("Content-Type")
		fmt.Println("contentType: ", contentType)
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
			fmt.Println("pload: ", pload)

		case strings.Contains(contentType, "x-www-form-urlencoded"):
			err := r.ParseForm()
			if err != nil {
				Error.Println(err)
			}
			pload.To = r.PostFormValue("to")

			//testmarkdown := r.PostFormValue("markdown")
			//if MsgType == "markdown"{
			//	var markdown Markdown
			//	err = json.Unmarshal([]byte(testmarkdown), &markdown)
			//	if err != nil{
			//		fmt.Println("json.Unmarshal err: ",err)
			//	}
			//	pload.Content = r.PostFormValue("content")
			//	pload.Markdown = markdown
			//}else {
			fmt.Println("test: ", MsgType)
			pload.Content = r.PostFormValue("content")
			fmt.Println("pload.Content: ", pload.Content)
			//}

		case strings.Contains(contentType, "markdown"):
			fmt.Println("r.Body: ", contentType)
			w.Header().Set("Content-Type", contentType)
			b, err := ioutil.ReadAll(r.Body)
			fmt.Println("r.Body: ", b)
			if err != nil {
				Error.Println(err)
			}
			err = pload.UnmarshalJSON(b)
			if err != nil {
				Error.Println(err)
			}

		default:
			fmt.Println("err")
			Error.Println("invalid Content-Type")
		}
		Info.Printf("#sendWechat# client:%s, to:%s, requestType:%s, content:%s\n", r.RemoteAddr, pload.To, contentType, pload.Content)
		// 判断发送内容是否是markdown格式的，默认非markdown格式
		if pload.MsgType == "markdown" {
			resp, err := dingding.SendMsgMarkdown("", pload.To,pload.Title,pload.Content)
			if err != nil {
				Error.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				Info.Printf(string(resp))
				w.Write(resp)
			}
		} else {
			resp, err := dingding.SendMsg("", pload.To, pload.Content)
			if err != nil {
				Error.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				Info.Printf(string(resp))
				w.Write(resp)
			}
		}

	})
	log.Fatal(http.ListenAndServe(httpAddr, nil))

}
