package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"echo-example/api/internal/handlers"
	"echo-example/api/internal/models"
	"echo-example/api/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status)
			return nil
		},
	}))

	e.Validator = &Validator{validator: validator.New()}

	db, err := openDB()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	service := services.New(models.NewModels(db))
	handler := handlers.New(service)

	v1 := e.Group("/v1")
	v1.GET("/employees", handler.EmployeeHandler.HandlerList)
	v1.GET("/employees/:id", handler.EmployeeHandler.HandlerGetByID)
	v1.POST("/employees", handler.EmployeeHandler.HandlerCreate)

	e.Logger.Fatal(e.Start(":8080"))
}

func openDB() (*sql.DB, error) {
	// TODO: Move this to a configuration struct
	db, err := sql.Open("postgres", "postgres://test-user:test-password@localhost:5432/test?sslmode=disable")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
