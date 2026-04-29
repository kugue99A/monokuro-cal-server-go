package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, eventHandler *EventHandler) {
	v1 := e.Group("/api/v1")

	events := v1.Group("/events")
	events.GET("", eventHandler.ListEvents)
	events.POST("", eventHandler.CreateEvent)
	events.GET("/:id", eventHandler.GetEvent)
	events.PUT("/:id", eventHandler.UpdateEvent)
	events.DELETE("/:id", eventHandler.DeleteEvent)
}
