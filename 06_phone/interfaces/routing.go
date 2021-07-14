package interfaces

import (
	"github.com/labstack/echo/v4"
	"phone/interfaces/handler"
)

type Router interface {
	Set(e *echo.Echo)
}

type router struct {
	handlers handler.AllHandler
}

func NewRouter(handlers handler.AllHandler) Router {
	return &router{handlers: handlers}
}

func (r router) Set(e *echo.Echo) {
	e.GET("/", r.handlers.GetAllEntriesHandler)
	e.PUT("/", r.handlers.SetNewEntryHandler)
}
