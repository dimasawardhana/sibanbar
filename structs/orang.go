package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Orang struct {
	gorm.Model
	Nama_lengkap  string    `json:"nama" form:"nama"`
	Alamat        string    `json:"alamat" form:"alamat"`
	Jenis_kelamin string    `json:"jenis_kelamin" form:"jenis_kelamin"`
	Tempat_lahir  string    `json:"tempat_lahir" form:"tempat_lahir"`
	Tanggal_lahir time.Time `json:"tanggal_lahir" form:"tanggal_lahir"`
	Status        string    `json:"status" form:"status"`
	Photo         string    `json:"photo" form:"photo"`
	Kelompok      Group     `gorm:"foreignKey:KelompokID;association_foreignkey:Id"`
	KelompokID    uint      `json:"kelompokId" form:"kelompokId"`
	Phone         string    `json:"phone" form:"phone"`
	Email         string    `json:"email" form:"email"`
}
