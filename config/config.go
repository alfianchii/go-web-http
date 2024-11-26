package config

import "fmt"

const (
	Port = 8080
	Host = "localhost"
)

var Address = fmt.Sprintf("%s:%d", Host, Port)