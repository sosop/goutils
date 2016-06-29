package goutils

import (
	"net/http"
)

// Middlerware http请求中间件
func Middlerware(handle http.Handler, f func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handle.ServeHTTP(w, r)
	})
}
