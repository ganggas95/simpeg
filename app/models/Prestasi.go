package models

import (
	"github.com/revel/revel"
)

type Prestasi struct {
	IdPrestasi    int64  `gorm:"column:id_prestasi;primary_key;AUTO_INCREMENT"`
	IdPegawai     int64  `gorm:"column:id_pegawai;index"`
	Prestasi      string `gorm:"column:prestasi;type:varchar(45)"`
	Penyelenggara string `gorm:"column:penyelenggara;type:varchar(45)"`
	Tahun         string `gorm:"column:tahun"`
}

func (prestasi *Prestasi) Validation(r *revel.Validation) {
	r.Check(prestasi.Prestasi, revel.Required{}, revel.MinSize{4}, revel.MaxSize{45})
	r.Check(prestasi.Penyelenggara, revel.Required{}, revel.MinSize{4})
	r.Check(prestasi.Tahun, revel.Required{}, revel.MinSize{4}, revel.MaxSize{4})
}
