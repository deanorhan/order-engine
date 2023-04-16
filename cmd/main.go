package main

import (
	"github.com/deanorhan/order-engine/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	api.Order(e)

	e.Logger.Fatal(e.Start(":8000"))
}
