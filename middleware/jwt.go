package middleware

import (
	"net/http"
	"web-http/config"
	"web-http/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
		jwtToken := session.Values["token"]

		if jwtToken == nil {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		
		_, err := utils.ValidateJWT(jwtToken.(string))
		if err != nil {
			utils.RemoveCookie(res, req, session)
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(res, req)
	})
}