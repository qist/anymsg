// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA8a797f8DecodeGithubComKexirongMsgSenderHttp(in *jlexer.Lexer, out *payload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "from":
			out.From = string(in.String())
		case "to":
			out.To = string(in.String())
		case "subject":
			out.Subject = string(in.String())
		case "content":
			out.Content = string(in.String())
		case "content_type":
			out.ContentType = string(in.String())
		case "msg_type":
			out.MsgType = string(in.String())
		case "title":
			out.Title = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA8a797f8EncodeGithubComKexirongMsgSenderHttp(out *jwriter.Writer, in payload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"from\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.From))
	}
	{
		const prefix string = ",\"to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.To))
	}
	{
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"content\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"content_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContentType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v payload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA8a797f8EncodeGithubComKexirongMsgSenderHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v payload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8a797f8EncodeGithubComKexirongMsgSenderHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *payload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA8a797f8DecodeGithubComKexirongMsgSenderHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *payload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8a797f8DecodeGithubComKexirongMsgSenderHttp(l, v)
}
