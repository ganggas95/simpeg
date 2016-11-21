package models

import (
	"github.com/revel/revel"
)

type Alamat struct {
	IdAlamat      int64  `gorm:"column:id_jabatan;primary_key;AUTO_INCREMENT"`
	IdPegawai     int64  `gorm:column:id_pegawai;index"`
	Dusun         string `gorm:"column:dusun;type:varchar(45)"`
	Desa          string `gorm:"column:desa;type:varchar(45)"`
	Kec           string `gorm:"column:kec;type:varchar(45)"`
	Kab           string `gorm:"column:kab;type:varchar(45)"`
	Prov          string `gorm:"column:prov;type:varchar(45)"`
	AlamatLengkap string `gorm:"column:alamat;type:varchar(255)"`
}

func (alamat *Alamat) Validation(r *revel.Validation) {
	r.Check(alamat.Dusun, revel.Required{}, revel.MaxSize{45})
	r.Check(alamat.Desa, revel.Required{}, revel.MaxSize{45})
	r.Check(alamat.Kec, revel.Required{}, revel.MaxSize{45})
	r.Check(alamat.Kab, revel.Required{}, revel.MaxSize{45})
	r.Check(alamat.Prov, revel.Required{}, revel.MaxSize{45})
}
