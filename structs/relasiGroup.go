package structs

import "github.com/jinzhu/gorm"

type relasiGrup struct {
	gorm.Model
	group         Group
	groupID       uint
	memberOf      Group
	memberID      uint
	relation_type string
}
