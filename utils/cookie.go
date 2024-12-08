package utils

import (
	"net/http"
	"strings"
	"web-http/config"

	"github.com/gorilla/sessions"
)

var (
	key = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

func Init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
	}
}

func CookieToString (res http.ResponseWriter) string {
	strCookie := ""
	for _, cookie := range res.Header()["Set-Cookie"] {
		if strings.Contains(cookie, config.SessionName) {
			strCookie = cookie
			break
		}
	}

	return strCookie
}