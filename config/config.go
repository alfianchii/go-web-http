package config

import "fmt"

const (
	Port = 3333
	Host = "localhost"
)

var Address = fmt.Sprintf("%s:%d", Host, Port)