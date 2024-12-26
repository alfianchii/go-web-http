package middleware

import (
	"net/http"
	"time"
	"web-http/config"
	"web-http/features/token"
	"web-http/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		session, err := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
		if err != nil || len(session.Values) == 0 {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}

		sessionExp, ok := session.Values["exp"].(int64)
		if !ok || time.Now().Unix() > sessionExp {
			utils.RemoveCookie(res, req, session)
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}

		sessionUsername := session.Values["username"]
		sessionToken := session.Values["token"]

		if sessionUsername == nil || sessionToken == nil {
			utils.RemoveCookie(res, req, session)
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		
		jwtToken, err := token.GetValidTokenFromUser(sessionUsername.(string))
		validToken := jwtToken.Value
		if err != nil {
			utils.RemoveCookie(res, req, session)
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		
		if validToken != sessionToken {
			utils.RemoveCookie(res, req, session)
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		
		_, err = utils.ValidateJWT(sessionToken.(string))
		if err != nil {
			utils.RemoveCookie(res, req, session)
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(res, req)
	})
}