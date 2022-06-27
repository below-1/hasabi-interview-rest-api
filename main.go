package main

import (
	"net/http"
	"rest-api/configs"
	"rest-api/controllers"
	"rest-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	app_configs := configs.GetConfig()
	db := configs.GetDb()

	// Run seeding in development
	if app_configs.AppEnv == "development" {
		db.AutoMigrate(&models.Peserta{}, &models.Soal{}, &models.JawabanPeserta{})
		// seed.Seed()
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/peserta", controllers.ListPeserta)
	e.POST("/api/peserta", controllers.CreatePeserta)
	e.DELETE("/api/peserta/:id", controllers.RemovePeserta)
	e.PUT("/api/peserta/:id", controllers.UpdatePeserta)
	e.GET("/api/soal", controllers.ListSoal)
	e.POST("/api/soal", controllers.CreateSoal)
	e.DELETE("/api/soal/:id", controllers.RemoveSoal)

	e.Logger.Fatal(e.Start(":1323"))
}
