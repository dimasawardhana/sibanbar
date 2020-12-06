package structs

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"size:255;unique;not null" json:"username" form:"username"`
	Password string `gorm:"size:255;not null" json:"password" form:"password"`
	Orang Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID uint `json:"orang_id" form:"orang_id"`
}