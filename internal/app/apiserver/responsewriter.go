package apiserver

import "net/http"

type responceWriter struct {
	http.ResponseWriter
	code int
}

func (w *responceWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
