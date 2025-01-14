// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

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

func easyjson6601e8cdDecodeTmpEasyjson(in *jlexer.Lexer, out *RawSettings) {
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
		case "denied_labels":
			if in.IsNull() {
				in.Skip()
				out.DeniedLabels = nil
			} else {
				in.Delim('[')
				if out.DeniedLabels == nil {
					if !in.IsDelim(']') {
						out.DeniedLabels = make([]string, 0, 4)
					} else {
						out.DeniedLabels = []string{}
					}
				} else {
					out.DeniedLabels = (out.DeniedLabels)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.DeniedLabels = append(out.DeniedLabels, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "mandatory_labels":
			if in.IsNull() {
				in.Skip()
				out.MandatoryLabels = nil
			} else {
				in.Delim('[')
				if out.MandatoryLabels == nil {
					if !in.IsDelim(']') {
						out.MandatoryLabels = make([]string, 0, 4)
					} else {
						out.MandatoryLabels = []string{}
					}
				} else {
					out.MandatoryLabels = (out.MandatoryLabels)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.MandatoryLabels = append(out.MandatoryLabels, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "constrained_labels":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.ConstrainedLabels = make(map[string]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 string
					v3 = string(in.String())
					(out.ConstrainedLabels)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjson6601e8cdEncodeTmpEasyjson(out *jwriter.Writer, in RawSettings) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"denied_labels\":"
		out.RawString(prefix[1:])
		if in.DeniedLabels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.DeniedLabels {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"mandatory_labels\":"
		out.RawString(prefix)
		if in.MandatoryLabels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.MandatoryLabels {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"constrained_labels\":"
		out.RawString(prefix)
		if in.ConstrainedLabels == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.ConstrainedLabels {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				out.String(string(v8Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RawSettings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeTmpEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RawSettings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeTmpEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RawSettings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeTmpEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RawSettings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeTmpEasyjson(l, v)
}
