package router

import (
	"server-go/operation"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Name     string
	Funcname echo.HandlerFunc
}

//Routers s
var Routers = []Router{
	{"/helloworld", operation.HelloWorld},
	{"/title", operation.Title},
}
