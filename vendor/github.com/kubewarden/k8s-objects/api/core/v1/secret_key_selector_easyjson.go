// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonE8b2103eDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *SecretKeySelector) {
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
		case "key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(string)
				}
				*out.Key = string(in.String())
			}
		case "name":
			out.Name = string(in.String())
		case "optional":
			out.Optional = bool(in.Bool())
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
func easyjsonE8b2103eEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in SecretKeySelector) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		if in.Key == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Key))
		}
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Optional {
		const prefix string = ",\"optional\":"
		out.RawString(prefix)
		out.Bool(bool(in.Optional))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecretKeySelector) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE8b2103eEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecretKeySelector) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE8b2103eEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecretKeySelector) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE8b2103eDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecretKeySelector) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE8b2103eDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
