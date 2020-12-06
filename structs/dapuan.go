package structs

import "github.com/jinzhu/gorm"

type Dapuan struct {
	gorm.Model
	Dapuan  string `json:"dapuan" form:"dapuan"`
	Orang   Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID uint `json:"orang_id" form:"orang_id"`
	Grup    Group `gorm:"foreignkey:GrupID;association_foreignkey:Id"`
	GrupID  uint `json:"group_id" form:"group_id"`
}
