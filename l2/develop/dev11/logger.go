package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Println(r.Method)
			switch r.Method {
			case http.MethodPost:
				logger.Println(r.Body)
			case http.MethodGet:
				logger.Println(r.URL)
			default:
				logger.Println("unknown method")
			}

			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Println(wrapped.status)
			fmt.Print(&buf)
		})
}
