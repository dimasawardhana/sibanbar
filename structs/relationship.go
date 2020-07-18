package structs

import "github.com/jinzhu/gorm"

type relationship struct {
	gorm.Model
	From          Orang
	idFrom        uint
	To            Orang
	idTo          uint
	Relation_type string
}
