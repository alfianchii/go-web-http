package middleware

import (
	"net/http"
)

func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if req.Method == http.MethodOptions {
				res.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(res, req)
    })
}