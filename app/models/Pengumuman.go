package models

import (
	"github.com/revel/revel"
	"time"
)

type Pengumuman struct {
	IdPengumuman int64     `gorm:"column:id_pengumuman;primary_key;AUTO_INCREMENT"`
	Judul        string    `gorm:"column:judul;type:varchar(45)"`
	Perihal      string    `gorm:"column:perihal;type:varchar(20)"`
	Tanggal      time.Time `gorm:"column:tanggal;type:date"`
	Isi          string    `gorm:"column:isi;type:text"`
	IdPegawai    int64     `gorm:column:id_pegawai;index"`
}

func (p *Pengumuman) Validation(r *revel.Validation) {
	r.Check(p.Judul, revel.Required{}, revel.MinSize{4}, revel.MaxSize{30})
	r.Check(p.Perihal, revel.Required{}, revel.MinSize{4}, revel.MaxSize{20})
	r.Check(p.Tanggal, revel.Required{})
	r.Check(p.Isi, revel.Required{})
}
