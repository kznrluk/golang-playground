package main

import (
	"github.com/labstack/echo/v4"
	"phone/interfaces"
	"phone/interfaces/handler"
)

func main() {
	e := echo.New()

	allHandler := handler.NewAllHandler()
	r := interfaces.NewRouter(allHandler)
	r.Set(e)

	e.Logger.Debug(e.Start(":8090"))
}
