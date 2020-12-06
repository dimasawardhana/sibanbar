package structs

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama"`
	Masjid       string `json:"masjid" form:"masjid"`
	Masjid_lang  float64 `json:"masjid_lang" form:"masjid_lang"`
	Masjid_lat   float64 `json:"masjid_lat" form:"masjid_lat"`
	Alamat       string `json:"alamat" form:"alamat"`
	Group_type   string `json:"group_type" form:"group_type"`
	Parent_group *Group `gorm:"foreignKey:parentId;association_foreignkey:Id" json:"parent_group" form:"parent_group"`
	ParentID     uint `json:"parent_id" form:"parent_id"`
}
