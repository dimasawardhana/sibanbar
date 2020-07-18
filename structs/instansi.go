package structs

import "github.com/jinzhu/gorm"

type Instansi struct{
	gorm.Model
	Nama_Instansi string
	Alamat_Instansi string
}