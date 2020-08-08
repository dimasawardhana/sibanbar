package structs

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	Username string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Orang Orang `gorm:"foreignkey:OrangID;association_foreignkey:Id"`
	OrangID uint
}