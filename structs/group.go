package structs

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Nama         string
	Masjid       string
	Masjid_lang  float64
	Masjid_lat   float64
	Alamat       string
	Group_type   string
	Parent_group *Group `gorm:"foreignKey:parentId;association_foreignkey:Id"`
	ParentID     uint
}
