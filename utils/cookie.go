package utils

import (
	"net/http"
	"web-http/config"

	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

func InitCookie() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int(config.TokenDuration.Seconds()),
		HttpOnly: true,
		Secure:   false,
	}
}

func RemoveCookie(res http.ResponseWriter, req *http.Request, session *sessions.Session) {
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	session.Save(req, res)
}