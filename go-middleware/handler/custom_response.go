package handler

import (
	"bytes"
	"net/http"
)

type customResponse struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func newCustomResponse(w http.ResponseWriter) *customResponse {
	return &customResponse{ResponseWriter: w, body: &bytes.Buffer{}}
}

func (rw *customResponse) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *customResponse) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}
