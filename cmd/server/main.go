package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kugue99A/monokuro-cal-server-go/internal/handler"
	"github.com/kugue99A/monokuro-cal-server-go/internal/infrastructure/postgres"
	eventusecase "github.com/kugue99A/monokuro-cal-server-go/internal/usecase/event"
	userusecase "github.com/kugue99A/monokuro-cal-server-go/internal/usecase/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "monokuro_cal"),
		getEnv("DB_SSLMODE", "disable"),
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := postgres.NewUserRepository(pool)
	eventRepo := postgres.NewEventRepository(pool)

	userUsecase := userusecase.NewUsecase(userRepo)
	eventUsecase := eventusecase.NewUsecase(eventRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	eventHandler := handler.NewEventHandler(eventUsecase)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.RegisterRoutes(e, userHandler, eventHandler)

	port := getEnv("PORT", "8080")
	e.Logger.Fatal(e.Start(":" + port))
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
