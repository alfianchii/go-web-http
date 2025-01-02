package middleware

import (
	"net/http"
	"web-http/utils"
)

func GuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		session, _ := utils.Store.Get(req, utils.SessionName)
		authHeader := req.Header.Get("Authorization")

		if authHeader != "" || session.Values["token"] != nil {
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(res, req)

	})
}