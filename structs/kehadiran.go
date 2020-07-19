package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Kehadiran struct {
	gorm.Model
	Jam_hadir  time.Time
	Keterangan string
	Event      Event `gorm:"foreignkey:EventID;association_foreignkey:Id"`
	EventID    uint
	Orang      Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID    uint
}
