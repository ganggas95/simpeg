package models

import (
	"github.com/revel/revel"
)

type Pendidikan struct {
	IdPendidikan int64  `gorm:"column:id_pendidikan;primary_key;AUTO_INCREMENT"`
	IdPegawai    int64  `gorm:column:id_pegawai;index"`
	Sekolah      string `gorm:column:sekolah;type:varchar(50)"`
	Tahun        string `gorm:column:tahun;`
}

type Tingkat struct {
	IdTingkat    int64  `gorm:"column:id_tingkat;primary_key;AUTO_INCREMENT"`
	IdPendidikan int64  `gorm:"column:id_pendidikan;index"`
	Tingkat      string `gorm:"column:tingkat;type:varchar(10)"`
}

func (p *Pendidikan) Validation(r *revel.Validation) {
	r.Check(p.Sekolah, revel.Required{}, revel.MinSize{5}, revel.MaxSize{40})
	r.Check(p.Tahun, revel.Required{}, revel.MinSize{4}, revel.MaxSize{4})
}

func (t *Tingkat) Validation(r *revel.Validation) {
	r.Check(t.Tingkat, revel.Required{})
}
