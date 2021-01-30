package operation

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func hellworld(c echo.Context) error {

	return c.String(http.StatusOK, "helloworld")
}
