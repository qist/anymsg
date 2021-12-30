// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package lark

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

func easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat(in *jlexer.Lexer, out *extend) {
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
		case "errcode":
			out.ErrCode = int64(in.Int64())
		case "errmsg":
			out.ErrMsg = string(in.String())
		case "access_token":
			out.AccToken = string(in.String())
		case "expires_in":
			out.TokenTS = int64(in.Int64())
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
func easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat(out *jwriter.Writer, in extend) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"errcode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ErrCode))
	}
	{
		const prefix string = ",\"errmsg\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ErrMsg))
	}
	{
		const prefix string = ",\"access_token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AccToken))
	}
	{
		const prefix string = ",\"expires_in\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TokenTS))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v extend) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v extend) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *extend) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *extend) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat(l, v)
}
func easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat1(in *jlexer.Lexer, out *JsonMsg) {
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

		case "msg_type":
			out.Msg_Type = string(in.String())


		case "text":
			easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat2(in, &out.Content)
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

func easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat1(out *jwriter.Writer, in JsonMsg) {
	out.RawByte('{')
	first := true
	_ = first


	{
		const prefix string = ",\"msg_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Msg_Type))
	}

	{
		const prefix string = ",\"content\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat2(out, in.Content)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JsonMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JsonMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JsonMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JsonMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat1(l, v)
}
func easyjson17a47f2aDecodeGithubComKexirongMsgSenderWechat2(in *jlexer.Lexer, out *Content) {
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
		case "text":
			out.Text = string(in.String())
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

func easyjson17a47f2aEncodeGithubComKexirongMsgSenderWechat2(out *jwriter.Writer, in Content) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	out.RawByte('}')
}
