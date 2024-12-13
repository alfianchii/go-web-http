package middleware

import (
	"net/http"
	"web-http/config"
	"web-http/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (res http.ResponseWriter, req *http.Request) {
		session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))

		_, ok := session.Values["username"].(string)
		if !ok {
			utils.SendResponse(res, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(res, req)
	})
}