package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	domain "github.com/kugue99A/monokuro-cal-server-go/internal/domain/event"
	eventusecase "github.com/kugue99A/monokuro-cal-server-go/internal/usecase/event"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	usecase *eventusecase.Usecase
}

func NewEventHandler(uc *eventusecase.Usecase) *EventHandler {
	return &EventHandler{usecase: uc}
}

type eventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}

type eventResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func toResponse(ev *domain.Event) eventResponse {
	return eventResponse{
		ID:          ev.ID.String(),
		Title:       ev.Title,
		Description: ev.Description,
		StartAt:     ev.StartAt,
		EndAt:       ev.EndAt,
		CreatedAt:   ev.CreatedAt,
		UpdatedAt:   ev.UpdatedAt,
	}
}

func (h *EventHandler) GetEvent(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	ev, err := h.usecase.GetEvent(c.Request().Context(), id)
	if err != nil {
		if err == domain.ErrNotFound {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, toResponse(ev))
}

func (h *EventHandler) ListEvents(c echo.Context) error {
	events, err := h.usecase.ListEvents(c.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}
	res := make([]eventResponse, len(events))
	for i, ev := range events {
		res[i] = toResponse(ev)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	var req eventRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	ev, err := h.usecase.CreateEvent(c.Request().Context(), eventusecase.CreateInput{
		Title:       req.Title,
		Description: req.Description,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
	})
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusCreated, toResponse(ev))
}

func (h *EventHandler) UpdateEvent(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	var req eventRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	ev, err := h.usecase.UpdateEvent(c.Request().Context(), id, eventusecase.UpdateInput{
		Title:       req.Title,
		Description: req.Description,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
	})
	if err != nil {
		if err == domain.ErrNotFound {
			return echo.ErrNotFound
		}
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, toResponse(ev))
}

func (h *EventHandler) DeleteEvent(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := h.usecase.DeleteEvent(c.Request().Context(), id); err != nil {
		if err == domain.ErrNotFound {
			return echo.ErrNotFound
		}
		return echo.ErrInternalServerError
	}
	return c.NoContent(http.StatusNoContent)
}
