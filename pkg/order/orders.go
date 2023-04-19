package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Order struct {
	Name string `json:"name"`
}

func Route(e *echo.Echo) {
	e.GET("/order", getOrders)
}

func getOrders(c echo.Context) error {
	order := Order{}

	return c.JSON(http.StatusOK, order)
}
