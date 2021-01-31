package endpoints

import (
	"example.com/someservice555/api/hello"
	"example.com/someservice555/api/welcome"
	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	Path     string
	Method   string
	Function gin.HandlerFunc
}

var Endpoints = []Endpoint{
	{
		Path: hello.Path,
		Method: hello.Method,
		Function: hello.ProcessRequest,
	},
	{
		Path: welcome.Path,
		Method: welcome.Method,
		Function: welcome.ProcessRequest,
	},
}
