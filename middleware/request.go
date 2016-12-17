package middleware

import "net/http"

type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func (ua *Request) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Content-Type", "application/json")
	next(rw, r)
}