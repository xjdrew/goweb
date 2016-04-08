package utils

import (
	"net/http"
)

type Middleware func(h http.Handler) http.Handler

func UseMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
