// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	_v5 "github.com/golang-jwt/jwt/v5"
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

func easyjsonB448b467DecodeGithubComArandichMarketplaceSdkModel(in *jlexer.Lexer, out *Claims) {
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
		case "user_id":
			out.UserID = uint64(in.Uint64())
		case "iss":
			out.Issuer = string(in.String())
		case "sub":
			out.Subject = string(in.String())
		case "aud":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Audience).UnmarshalJSON(data))
			}
		case "exp":
			if in.IsNull() {
				in.Skip()
				out.ExpiresAt = nil
			} else {
				if out.ExpiresAt == nil {
					out.ExpiresAt = new(_v5.NumericDate)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ExpiresAt).UnmarshalJSON(data))
				}
			}
		case "nbf":
			if in.IsNull() {
				in.Skip()
				out.NotBefore = nil
			} else {
				if out.NotBefore == nil {
					out.NotBefore = new(_v5.NumericDate)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.NotBefore).UnmarshalJSON(data))
				}
			}
		case "iat":
			if in.IsNull() {
				in.Skip()
				out.IssuedAt = nil
			} else {
				if out.IssuedAt == nil {
					out.IssuedAt = new(_v5.NumericDate)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.IssuedAt).UnmarshalJSON(data))
				}
			}
		case "jti":
			out.ID = string(in.String())
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
func easyjsonB448b467EncodeGithubComArandichMarketplaceSdkModel(out *jwriter.Writer, in Claims) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"user_id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.UserID))
	}
	if in.Issuer != "" {
		const prefix string = ",\"iss\":"
		out.RawString(prefix)
		out.String(string(in.Issuer))
	}
	if in.Subject != "" {
		const prefix string = ",\"sub\":"
		out.RawString(prefix)
		out.String(string(in.Subject))
	}
	if len(in.Audience) != 0 {
		const prefix string = ",\"aud\":"
		out.RawString(prefix)
		out.Raw((in.Audience).MarshalJSON())
	}
	if in.ExpiresAt != nil {
		const prefix string = ",\"exp\":"
		out.RawString(prefix)
		out.Raw((*in.ExpiresAt).MarshalJSON())
	}
	if in.NotBefore != nil {
		const prefix string = ",\"nbf\":"
		out.RawString(prefix)
		out.Raw((*in.NotBefore).MarshalJSON())
	}
	if in.IssuedAt != nil {
		const prefix string = ",\"iat\":"
		out.RawString(prefix)
		out.Raw((*in.IssuedAt).MarshalJSON())
	}
	if in.ID != "" {
		const prefix string = ",\"jti\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Claims) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB448b467EncodeGithubComArandichMarketplaceSdkModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Claims) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB448b467EncodeGithubComArandichMarketplaceSdkModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Claims) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB448b467DecodeGithubComArandichMarketplaceSdkModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Claims) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB448b467DecodeGithubComArandichMarketplaceSdkModel(l, v)
}
