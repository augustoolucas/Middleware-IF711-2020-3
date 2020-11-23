package serverproxy

import "Middleware-IF711-2020-3/l5/auxiliar"

type ServerProxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

type Invocation struct {
	Host     string
	Port     int
	Response auxiliar.Response
}
