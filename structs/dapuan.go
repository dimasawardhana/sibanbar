package structs

import "github.com/jinzhu/gorm"

type Dapuan struct {
	gorm.Model
	Dapuan  string
	Orang   Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID uint
	Grup    Group `gorm:"foreignkey:GrupID;association_foreignkey:Id"`
	GrupID  uint
}
