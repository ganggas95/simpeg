package models

import (
	"github.com/revel/revel"
)

type Jabatan struct {
	IdJabatan int64  `gorm:"column:id_jabatan;primary_key;AUTO_INCREMENT"`
	IdPegawai int64  `gorm:column:id_pegawai;index"`
	Jabatan   string `gorm:"column:jabatan;type:varchar(45)"`
}

func (jabatan *Jabatan) Validaton(r *revel.Validation) {
	r.Check(jabatan.Jabatan, revel.Required{}, revel.MaxSize{20})
}
