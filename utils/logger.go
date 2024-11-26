package utils

import (
	"fmt"
	"net/http"
)

func LogRequest(r *http.Request) {
	fmt.Printf("Request Method: %s\n", r.Method)
	fmt.Printf("Request URL: %s\n", r.URL)
	fmt.Printf("Request Path: %s\n", r.URL.Path)
	fmt.Printf("Request Header: %v\n", r.Header)
	fmt.Printf("Request Body: %v\n", r.Body)
	fmt.Printf("======================================" + "\n")
}