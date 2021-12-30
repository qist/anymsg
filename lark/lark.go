package lark

import (
	"bytes"
	//"encoding/gob"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	AccTokenURL       = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	SendmsgURL        = " https://open.feishu.cn/open-apis/bot/v2/hook/%s"
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



type Lark struct {
	SendmsgURL string `json:"sendmsg_url"`
	Token      string `json:"token"`
}

//easyjson:json
type Content struct {
	Text string `json:"text"`
}
func New(url, token string) *Lark {
	return &Lark{
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
	Receive_Id      string      `json:"receive_id"`
	Msg_Type string  `json:"msg_type"`
	Content    Content `json:"content"`
}
//type Post struct {
//	Title string `json:"title"`
//	Content []interface{} `json:"text"`
//}
//// 定义PostMsg格式的消息体
////easyjson:json
//type PostMsg struct {
//	Msg_Type string `json:"msg_type"`
//	Post Post `json:"post"`
//}

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

func (lark Lark) SendMsg(content string) ([]byte, error) {

	msg := JsonMsg{
		Msg_Type: "text",
		Content: Content{
			Text: content,
		},
	}
	//msg := Content{
	//	Text: content,
	//}
    fmt.Println("msg: ",msg)

	jmsg, err := msg.MarshalJSON()
	if err != nil {
		return nil, err
	}
	fmt.Println("jmsg: ",string(jmsg))

	postSendmsgURL := fmt.Sprintf(lark.SendmsgURL, lark.Token)
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
//func (lark Lark) SendMsgPost(atMobiles, atUserIds,title string, content []) ([]byte, error) {
//	msg := PostMsg{
//
//		Msg_Type: "post",
//
//		Post: Post{
//			Title: title,
//			Content: content,
//		},
//	}
//
//	msgjson, err := json.Marshal(msg)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Println(dingtalk.SendmsgURL)
//	fmt.Println(dingtalk.Token)
//	postSendmsgURL := fmt.Sprintf(dingtalk.SendmsgURL, dingtalk.Token)
//	fmt.Println(postSendmsgURL)
//	rsp, err := http.Post(postSendmsgURL, "application/json;charset=utf-8", bytes.NewReader(msgjson))
//	if err != nil {
//		return nil, err
//	}
//	defer rsp.Body.Close()
//	return ioutil.ReadAll(rsp.Body)
//
//	return nil, nil
//}
