package utils

import "strings"

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func CapitalizeFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}