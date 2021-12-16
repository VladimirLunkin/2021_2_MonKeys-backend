// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson8ceb9162DecodeDripappInternalDripappModels(in *jlexer.Lexer, out *Photo) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "photo":
			out.Path = string(in.String())
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
func easyjson8ceb9162EncodeDripappInternalDripappModels(out *jwriter.Writer, in Photo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"photo\":"
		out.RawString(prefix[1:])
		out.String(string(in.Path))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Photo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ceb9162EncodeDripappInternalDripappModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Photo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ceb9162EncodeDripappInternalDripappModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Photo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ceb9162DecodeDripappInternalDripappModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Photo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ceb9162DecodeDripappInternalDripappModels(l, v)
}
