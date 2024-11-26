package middleware

import (
	"net/http"
	"web-http/utils"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		utils.LogRequest(req)
		next.ServeHTTP(res, req)
	})
}