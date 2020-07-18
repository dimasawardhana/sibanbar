package structs

import "github.com/jinzhu/gorm"
import "time"

type Kehadiran struct{
	gorm.Model
	Jam_hadir time.Time
	Keterangan int
	Event Event `gorm:"foreignkey:EventID;association_foreignkey:Id"`
	EventID uint
	Orang Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID uint
}