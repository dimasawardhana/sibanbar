package structs

import "github.com/jinzhu/gorm"

type Relationship struct {
	gorm.Model
	From          Orang
	IdFrom        uint	`json:"from" form:"from"`
	To            Orang
	IdTo          uint	`json:"to" form:"to"`
	Relation_type string `json:"relation_type" form:"relation_type"`
}
