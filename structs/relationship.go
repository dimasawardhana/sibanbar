package structs

import "github.com/jinzhu/gorm"

type Relationship struct {
	gorm.Model
	From          Orang
	IdFrom        uint
	To            Orang
	IdTo          uint
	Relation_type string
}
