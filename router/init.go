package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ojbk struct {
	Ok     string `json:"ok"`
	Status int    `status:"status"`
}

func Listen(port string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	for _, v := range Routers {
		e.GET(v.Name, v.Funcname)
	}

	e.Logger.Fatal(e.Start(port))

}

func hello(c echo.Context) error {
	i := ojbk{
		Ok:     "ojbk",
		Status: 8848,
	}
	return c.JSON(http.StatusOK, i)
	// return c.String(http.StatusOK, "Hello, World!")
}
