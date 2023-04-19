package main

import (
	"github.com/deanorhan/order-engine/pkg/order"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	order.Route(e)

	e.Logger.Fatal(e.Start(":8000"))
}
