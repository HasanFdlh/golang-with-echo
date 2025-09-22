// internal/handler/user_handler.go
package handler

import (
	"ms-golang-echo/config"
	"ms-golang-echo/internal/model"
	"ms-golang-echo/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// Create
func (h *UserHandler) Create(c echo.Context) error {
	var req model.UserRequest
	if err := c.Bind(&req); err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, "invalid request", nil)
	}

	if err := c.Validate(&req); err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	user := model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.usecase.Create(&user); err != nil {
		return config.JSONResponse(c, http.StatusInternalServerError, "failed to create user", nil)
	}

	user.CreatedAtStr = model.JSONTime(user.CreatedAt)
	user.UpdatedAtStr = model.JSONTime(user.UpdatedAt)

	return config.Success(c, user)
}

// Get by ID
func (h *UserHandler) GetByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, "invalid user id", nil)
	}

	user, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return config.JSONResponse(c, http.StatusNotFound, "user not found", nil)
	}

	user.CreatedAtStr = model.JSONTime(user.CreatedAt)
	user.UpdatedAtStr = model.JSONTime(user.UpdatedAt)

	return config.Success(c, user)
}

// List
func (h *UserHandler) List(c echo.Context) error {
	users, err := h.usecase.List()
	if err != nil {
		return config.JSONResponse(c, http.StatusInternalServerError, "failed to get users", nil)
	}

	// Format waktu
	for i := range users {
		users[i].CreatedAtStr = model.JSONTime(users[i].CreatedAt)
		users[i].UpdatedAtStr = model.JSONTime(users[i].UpdatedAt)
	}

	return config.Success(c, users)
}

// Update
func (h *UserHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, "invalid user id", nil)
	}

	// cek user ada
	user, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return config.JSONResponse(c, http.StatusNotFound, "user not found", nil)
	}

	// bind ke request
	var req model.UserRequest
	if err := c.Bind(&req); err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, "invalid request body", nil)
	}

	// validasi
	if err := c.Validate(&req); err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	// update field
	user.Name = req.Name
	user.Email = req.Email

	if err := h.usecase.Update(user); err != nil {
		return config.JSONResponse(c, http.StatusInternalServerError, "failed to update user", nil)
	}

	user.CreatedAtStr = model.JSONTime(user.CreatedAt)
	user.UpdatedAtStr = model.JSONTime(user.UpdatedAt)

	return config.Success(c, user)
}

// Delete
func (h *UserHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return config.JSONResponse(c, http.StatusBadRequest, "invalid user id", nil)
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		return config.JSONResponse(c, http.StatusInternalServerError, "failed to delete user", nil)
	}

	return config.JSONResponse(c, http.StatusOK, "user deleted", nil)
}
