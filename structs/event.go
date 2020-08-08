package structs

import "github.com/jinzhu/gorm"
import "time"

type Event struct{
	gorm.Model
	Nama string
	Tanggal time.Time
	Tempat string
	Tempat_lat float64
	Tempat_lang float64
	Poster string
}