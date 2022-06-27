package controllers

import (
	"net/http"
	"rest-api/configs"
	"rest-api/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CreateSoalInput struct {
	Pertanyaan string `json:"pertanyaan"`
	Aktif      bool   `json:"aktif"`
	Kode       string `json:"kode"`
	Point      int    `json:"point"`
}

type UpdateSoalInput struct {
	Pertanyaan string `json:"pertanyaan"`
	Aktif      bool   `json:"aktif"`
	Kode       string `json:"kode"`
	Point      int    `json:"point"`
}

func CreateSoal(c echo.Context) (err error) {
	var payload CreateSoalInput
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	if err = c.Validate(payload); err != nil {
		return echo.ErrBadRequest
	}

	db := configs.GetDb()
	err = db.Create(&models.Soal{
		Pertanyaan: payload.Pertanyaan,
		Aktif:      payload.Aktif,
		Kode:       payload.Kode,
		Point:      uint(payload.Point),
	}).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"status": "OK",
	})
}

func ListSoal(c echo.Context) (err error) {
	db := configs.GetDb()
	var items []models.Soal
	if err = db.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, items)
}

func RemoveSoal(c echo.Context) (err error) {
	idStr := c.Param("id")
	var id int
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	db := configs.GetDb()
	err = db.Delete(&models.Soal{
		ID: id,
	}).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"status": "OK",
	})
}
