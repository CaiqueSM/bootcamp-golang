package web

import (
	"net/http"
	"strconv"
)

type Response struct {
	Code  string      `jason:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewRespponse(code int, data interface{}, err string)Response {
	if code < http.StatusMultipleChoices {
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	return Response{strconv.FormatInt(int64(code), 10), nil, err}
}
