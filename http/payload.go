package http

//easyjson:json
type payload struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}
