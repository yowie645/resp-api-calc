package main

import (
	"log"
	calcservice "resp-api-calc/internal/calcService"
	"resp-api-calc/internal/db"
	"resp-api-calc/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	e := echo.New()

	calcRepo := calcservice.NewCalculationsRepository(database)
	calcService := calcservice.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}
