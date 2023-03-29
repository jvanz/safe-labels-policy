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

func easyjson18107e85DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *SELinuxOptions) {
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
		case "level":
			out.Level = string(in.String())
		case "role":
			out.Role = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "user":
			out.User = string(in.String())
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
func easyjson18107e85EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in SELinuxOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Level != "" {
		const prefix string = ",\"level\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Level))
	}
	if in.Role != "" {
		const prefix string = ",\"role\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Role))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SELinuxOptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18107e85EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SELinuxOptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18107e85EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SELinuxOptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18107e85DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SELinuxOptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18107e85DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
