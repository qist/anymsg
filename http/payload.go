package http

//easyjson:json
type payload struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	MsgType     string `json:"msg_type"`
	Title       string `json:"title"`
}

//func (conf *Payload) String() string {
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
