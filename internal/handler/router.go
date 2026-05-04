package handler

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, userHandler *UserHandler, eventHandler *EventHandler) {
	v1 := e.Group("/api/v1")

	users := v1.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.GET("/:id", userHandler.GetUser)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)

	// イベントはユーザー配下と単体の両方
	users.GET("/:user_id/events", eventHandler.ListEvents)
	users.POST("/:user_id/events", eventHandler.CreateEvent)

	events := v1.Group("/events")
	events.GET("/:id", eventHandler.GetEvent)
	events.PUT("/:id", eventHandler.UpdateEvent)
	events.DELETE("/:id", eventHandler.DeleteEvent)
}
