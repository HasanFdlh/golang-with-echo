package routes

import (
	"ms-golang-echo/internal/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group, h *handler.UserHandler) {
	e.POST("", h.Create)
	e.GET("/:id", h.GetByID)
	e.GET("", h.List)
	e.PUT("/:id", h.Update)
	e.DELETE("/:id", h.Delete)
}
