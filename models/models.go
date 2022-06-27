package models

import (
	"time"

	"gorm.io/gorm"
)

type Peserta struct {
	gorm.Model
	ID           int       `json:"id" gorm:"primaryKey"`
	Nama         string    `json:"nama"`
	JenisKelamin string    `json:"jenis_kelamin"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	NoHp         string    `json:"no_hp" gorm:"uniqueIndex"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
}

type Soal struct {
	gorm.Model
	ID         int    `json:"id" gorm:"primaryKey"`
	Pertanyaan string `json:"pertanyaan"`
	Aktif      bool   `json:"aktif"`
	Kode       string `json:"kode" gorm:"uniqueIndex"`
	Point      uint   `json:"point"`
}

type JawabanPeserta struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	SoalID    int    `json:"soal_id"`
	Soal      Soal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PesertaID int    `json:"peserta_id"`
	Peserta   Soal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Jawaban   string `json:"jawaban"`
}
