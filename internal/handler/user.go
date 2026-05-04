package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	domain "github.com/kugue99A/monokuro-cal-server-go/internal/domain/user"
	userusecase "github.com/kugue99A/monokuro-cal-server-go/internal/usecase/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase *userusecase.Usecase
}

func NewUserHandler(uc *userusecase.Usecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

type userCreateRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type userUpdateRequest struct {
	Name string `json:"name"`
}

type userResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toUserResponse(u *domain.User) userResponse {
	return userResponse{
		ID:        u.ID.String(),
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	user, err := h.usecase.GetUser(c.Request().Context(), id)
	if err != nil {
		if err == domain.ErrNotFound {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, toUserResponse(user))
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req userCreateRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	user, err := h.usecase.CreateUser(c.Request().Context(), userusecase.CreateInput{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, toUserResponse(user))
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	var req userUpdateRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	user, err := h.usecase.UpdateUser(c.Request().Context(), id, userusecase.UpdateInput{
		Name: req.Name,
	})
	if err != nil {
		if err == domain.ErrNotFound {
			return echo.ErrNotFound
		}
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, toUserResponse(user))
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := h.usecase.DeleteUser(c.Request().Context(), id); err != nil {
		if err == domain.ErrNotFound {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return c.NoContent(http.StatusNoContent)
}
