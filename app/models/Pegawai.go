package models

import (
	"github.com/revel/revel"
	"time"
)

type Pegawai struct {
	IdPegawai    int64        `gorm:"column:id_pegawai;primary_key;AUTO_INCREMENT"`
	Nama         string       `gorm:"column:nama_pegawai;type:varchar(45)"`
	TempatLahir  string       `gorm:"column:tempat_lahir;type:varchar(45)"`
	TanggalLahir time.Time    `gorm:"column:tanggal_lahir;type:date"`
	Agama        string       `gorm:"column:agama;type:varchar(45)"`
	Jabatan      Jabatan      `gorm:"ForeignKey:IdPegawai"`
	Alamat       Alamat       `gorm:"ForeignKey:IdPegawai"`
	Pendidikan   []Pendidikan `gorm:"ForeignKey:IdPegawai"`
	Prestasi     []Prestasi   `gorm:"ForeignKey:IdPegawai"`
	IdUser       int64        `gorm:"column:id_user;index"`
}

func (pegawai *Pegawai) Validation(r *revel.Validation) {
	r.Check(pegawai.Nama, revel.Required{}, revel.MinSize{3})
	r.Check(pegawai.TempatLahir, revel.Required{}, revel.MinSize{3})
	r.Check(pegawai.TanggalLahir, revel.Required{})
	r.Check(pegawai.Agama, revel.Required{}, revel.MinSize{3}, revel.MaxSize{15})
}
