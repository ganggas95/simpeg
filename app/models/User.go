package models

import (
	"github.com/revel/revel"
	"regexp"
)

type User struct {
	IdUser   int64   `gorm:"column:id_user;primary_key;AUTO_INCREMENT"`
	Pegawai  Pegawai `gorm:"ForeignKey:IdUser"`
	Email    string  `gorm:"column:email;type:varchar(45)"`
	Username string  `gorm:"column:username;type:varchar(20)"`
	Password []byte  `gorm:"column:password"`
}

var emailregex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (user *User) Validasi(r *revel.Validation) {
	r.Check(user.Username, revel.Required{}, revel.MinSize{4}, revel.MaxSize{15})
	r.Check(user.Password, revel.Required{})
	r.Check(user.Email, revel.Match{emailregex})
}
