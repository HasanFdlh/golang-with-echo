package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func Success(c echo.Context, data interface{}) error {
	return JSONResponse(c, http.StatusOK, "success", data)
}

func Error(c echo.Context, message string) error {
	return JSONResponse(c, http.StatusBadRequest, message, nil)
}
