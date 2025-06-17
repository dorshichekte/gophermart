package middleware

import (
	"io"
	"net/http"
)

type middleware = func(f http.Handler) http.Handler

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

type ContextKey string
