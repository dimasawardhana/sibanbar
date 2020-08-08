package structs

import "github.com/jinzhu/gorm"

type Kesibukan struct {
	gorm.Model
	Orang      Orang
	OrangID    uint
	Instansi   Instansi
	InstansiID uint
	Status     string
	Kedudukan  string
}
