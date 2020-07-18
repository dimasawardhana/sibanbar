package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Orang struct {
	gorm.Model
	Nama_lengkap  string
	Alamat        string
	Tempat_lahir  string
	Tanggal_lahir time.Time
	Status        string
	Photo         string
	Kelompok      Group `gorm:"foreignKey:KelompokID;association_foreignkey:Id"`
	KelompokID    uint
}
