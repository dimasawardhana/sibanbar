package structs

import "github.com/jinzhu/gorm"

type Instansi struct{
	gorm.Model
	Nama_Instansi string `json:"nama_instansi" form:"nama_instansi"`
	Alamat_Instansi string `json:"alamat_instansi" form:"alamat_instansi"`
}