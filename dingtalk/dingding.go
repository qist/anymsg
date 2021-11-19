package dingtalk

import (
	"bytes"

	"encoding/json"

	//"encoding/gob"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	AccTokenURL       = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	SendmsgURL        = "https://oapi.dingtalk.com/robot/send?access_token=%s"
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



type Dingtalk struct {
	SendmsgURL string `json:"sendmsg_url"`
	Token      string `json:"token"`
}

type Content struct {
	Content string `json:"content"`
}

func New(url, token string) *Dingtalk {
	return &Dingtalk{
		SendmsgURL: url,
		Token:      token,
	}
}

type At struct {
	AtMobiles string `json:"atMobiles"`
	AtUserIds string `json:"atUserIds"`
	IsAtAll   bool   `json:"isAtAll"`
}

//easyjson:json
type JsonMsg struct {
	At      At      `json:"at"`
	MsgType string  `json:"msgtype"`
	Text    Content `json:"text"`
}
type Markdown struct {
	Title string `json:"title"`
	Text string `json:"text"`
}
// 定义Markdown格式的消息体
//easyjson:json
type MarkdownMsg struct {
	At      At      `json:"at"`
	MsgType string `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

/*func checkErr(err error){

}*/

//func (conf *MarkdownMsg) String() string {
//	b, err := json.Marshal(*conf)
//	if err != nil {
//		return fmt.Sprintf("%+v", *conf)
//	}
//	var out bytes.Buffer
//	err = json.Indent(&out, b, "", "    ")
//	if err != nil {
//		return fmt.Sprintf("%+v", *conf)
//	}
//	return out.String()
//}

func (dingtalk Dingtalk) SendMsg(atMobiles, atUserIds, content string) ([]byte, error) {

	msg := JsonMsg{
		At: At{
			AtMobiles: atMobiles,
			AtUserIds: atUserIds,
			IsAtAll:   false,
		},
		MsgType: "text",
		Text: Content{
			Content: content,
		},
	}

	jmsg, err := msg.MarshalJSON()
	if err != nil {
		return nil, err
	}

	postSendmsgURL := fmt.Sprintf(dingtalk.SendmsgURL, dingtalk.Token)
	fmt.Println("url: ",postSendmsgURL)
	rsp, err := http.Post(postSendmsgURL, "application/json", bytes.NewBuffer(jmsg))
	//rsp, err := TLSClient.Post(postSendmsgURL, "application/json;charset=utf-8", bytes.NewReader(jmsg))
	if err != nil {
		fmt.Println("TLSClient.Post: ",err)
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)

	//return nil, err
}

// 发送dingtalkMarkdown格式的方法
func (dingtalk Dingtalk) SendMsgMarkdown(atMobiles, atUserIds,title, content string) ([]byte, error) {
	msg := MarkdownMsg{
		At: At{
			AtMobiles: atMobiles,
			AtUserIds: atUserIds,
			IsAtAll:   false,
		},
		MsgType: "markdown",

		Markdown: Markdown{
			Title: title,
			Text: content,
		},
	}

	msgjson, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	fmt.Println(dingtalk.SendmsgURL)
	fmt.Println(dingtalk.Token)
	postSendmsgURL := fmt.Sprintf(dingtalk.SendmsgURL, dingtalk.Token)
	fmt.Println(postSendmsgURL)
	rsp, err := http.Post(postSendmsgURL, "application/json;charset=utf-8", bytes.NewReader(msgjson))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)

	return nil, nil
}
