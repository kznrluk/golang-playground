package handler

import (
	"github.com/labstack/echo/v4"
	"phone/domain"
	"phone/usecase"
)

type AllHandler struct {
	GetAllEntriesHandler echo.HandlerFunc
	SetNewEntryHandler   echo.HandlerFunc
}

func NewAllHandler() AllHandler {
	return AllHandler{
		GetAllEntriesHandler: ProvideGetAllEntriesHandler(),
		SetNewEntryHandler:   SetNewEntryHandler(),
	}
}

func ProvideGetAllEntriesHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		uc := usecase.NewEntry()

		entries, err := uc.GetAllEntries()
		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, entries)
	}
}

func SetNewEntryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		uc := usecase.NewEntry()

		var put domain.Entry
		err := c.Bind(&put)
		if err != nil {
			return c.String(500, err.Error())
		}

		err = uc.AddEntry(put.Id, put.Name, put.Phone)
		if err != nil {
			return c.String(500, err.Error())
		}

		return c.String(200, "OK")
	}
}
