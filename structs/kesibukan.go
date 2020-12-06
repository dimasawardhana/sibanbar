package structs

import "github.com/jinzhu/gorm"

type Kesibukan struct {
	gorm.Model
	Orang      Orang
	OrangID    uint `json:"orang_id" form:"orang_id"`
	Instansi   Instansi
	InstansiID uint `json:"instansi_id" form:"instansi_id"`
	Status     string `json:"status" form:"status"`
	Kedudukan  string `json:"kedudukan" form:"kedudukan"`
}
