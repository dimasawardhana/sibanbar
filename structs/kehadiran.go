package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Kehadiran struct {
	gorm.Model
	Jam_hadir  time.Time `json:"jam_hadir" form:"jam_hadir"`
	Keterangan string `json:"keterangan" form:"keterangan"`
	Event      Event `gorm:"foreignkey:EventID;association_foreignkey:Id"`
	EventID    uint `json:"event_id" form:"event_id"`
	Orang      Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID    uint `json:"orang_id" form:"orang_id"`
}
