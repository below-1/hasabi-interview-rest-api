package controllers

import (
	"net/http"
	"rest-api/configs"
	"rest-api/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type CreatePesertaInput struct {
	Nama         string    `json:"nama"`
	JenisKelamin string    `json:"jenis_kelamin"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	NoHp         string    `json:"no_hp" gorm:"uniqueIndex"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
}

func CreatePeserta(c echo.Context) (err error) {
	var payload CreatePesertaInput
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	if err = c.Validate(payload); err != nil {
		return echo.ErrBadRequest
	}

	db := configs.GetDb()
	err = db.Create(&models.Peserta{
		Nama:         payload.Nama,
		JenisKelamin: payload.JenisKelamin,
		TempatLahir:  payload.TempatLahir,
		TanggalLahir: payload.TanggalLahir,
		NoHp:         payload.NoHp,
		Email:        payload.Email,
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

func ListPeserta(c echo.Context) (err error) {
	db := configs.GetDb()
	var items []models.Peserta
	if err = db.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, items)
}

func RemovePeserta(c echo.Context) (err error) {
	idStr := c.Param("id")
	var id int
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	db := configs.GetDb()
	err = db.Delete(&models.Peserta{
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

type UpdatePesertaInput struct {
	Nama         string    `json:"nama"`
	JenisKelamin string    `json:"jenis_kelamin"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	NoHp         string    `json:"no_hp" gorm:"uniqueIndex"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
}

func UpdatePeserta(c echo.Context) (err error) {
	idString := c.Param("id")
	var id int
	id, err = strconv.Atoi(idString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	var input UpdatePesertaInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	if err = c.Validate(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	db := configs.GetDb()
	var peserta models.Peserta
	if err = db.Where("id = ?", id).First(&peserta).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	peserta.Nama = input.Nama
	peserta.TanggalLahir = input.TanggalLahir
	peserta.NoHp = input.NoHp
	peserta.Email = input.Email
	peserta.TempatLahir = input.TempatLahir

	if err = db.Save(&peserta).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"status": "OK",
	})
}
