package structs

import "github.com/jinzhu/gorm"
import "time"

type Event struct{
	gorm.Model
	Nama string `json:"nama" form:"nama"`
	Tanggal time.Time `json:"tanggal" form:"tanggal"`
	Tempat string `json:"tempat" form:"tempat"`
	Tempat_lat float64 `json:"tempat_lat" form:"tempat_lat"`
	Tempat_lang float64 `json:"tempat_lang" form:"tempat_lang"`
	Poster string `json:"poster" form:"poster"`
}