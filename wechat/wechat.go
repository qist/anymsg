package wechat

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mailru/easyjson"
)

const (
	AccTokenURL       = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	SendmsgURL        = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
	TokeExpSec  int64 = 7200
)

/*
   "errcode": 0，
   "errmsg": "ok"，
   "access_token": "accesstoken000001",
   "expires_in": 7200
*/

//easyjson:json
type extend struct {
	ErrCode  int64  `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	AccToken string `json:"access_token"`
	TokenTS  int64  `json:"expires_in"`
}

type WeChat struct {
	CorpID  string
	AgentID int
	Secret  string
	*extend
}

var TLSClient *http.Client

func init() {
	/* file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	   if err != nil {
	       log.Fatalln("LogFile ioError:", err)
	   } */
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM([]byte(wxRootPEM))
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}

	TLSClient = &http.Client{Transport: tr}

}

func New(CorpID string, AgentID int, Secret string) *WeChat {
	return &WeChat{
		CorpID:  CorpID,
		AgentID: AgentID,
		Secret:  Secret,
		extend:  &extend{},
	}
}

type Content struct {
	Content string `json:"content"`
}

//easyjson:json
type JsonMsg struct {
	ToUser  string  `json:"touser,omitempty"`
	ToParty string  `json:"toparty,omitempty"`
	MsgType string  `json:"msgtype"`
	AgentID int     `json:"agentid"`
	Text    Content `json:"text"`
}

/*func checkErr(err error){

}*/

func (wx *WeChat) GetAccToken() error {
	getAccTokenURL := fmt.Sprintf(AccTokenURL, wx.CorpID, wx.Secret)

	rsp, err := TLSClient.Get(getAccTokenURL)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	err = easyjson.UnmarshalFromReader(rsp.Body, wx.extend)

	if err != nil {
		return err
	}
	//errcode := json.Get("errcode").MustInt(1)
	if wx.ErrCode != 0 {
		return fmt.Errorf("get WeChat Access Token error: %s", wx.ErrMsg)
	}
	wx.TokenTS += time.Now().Unix()
	return nil
}

func (wx WeChat) SendMsg(touser, toparty, content string) ([]byte, error) {

	msg := JsonMsg{
		ToUser:  touser,
		ToParty: toparty,
		MsgType: "text",
		AgentID: wx.AgentID,
		Text: Content{
			Content: content,
		},
	}

	if wx.AccToken == "" || wx.TokenTS-time.Now().Unix() <= 0 {
		var err error
		for i := 0; i < 3; i++ {
			err = wx.GetAccToken()
			if err == nil {
				break
			}
		}
		if err != nil {
			return nil, err
		}
	}

	jmsg, err := msg.MarshalJSON()
	if err != nil {
		return nil, err
	}

	postSendmsgURL := fmt.Sprintf(SendmsgURL, wx.AccToken)
	rsp, err := TLSClient.Post(postSendmsgURL, "application/json;charset=utf-8", bytes.NewReader(jmsg))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)

	//return nil, err
}
