package structs

import "github.com/jinzhu/gorm"

type Relationship struct {
	gorm.Model
	From          Orang `json:"from"`
	IdFrom        uint	`json:"from_id" form:"from_id"`
	To            Orang `json:"to" form:"to"`
	IdTo          uint	`json:"to_id" form:"to_id"`
	Relation_type string `json:"relation_type" form:"relation_type"`
}
