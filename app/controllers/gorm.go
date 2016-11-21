package controllers

import (
	"github.com/ganggas95/simpeg/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Database struct {
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (d Database) initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:keyfaton@/db_proj?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		CheckError(err)
	}
	d.createTable(db)
	return db
}
func (d Database) createTable(db *gorm.DB) {
	db.SingularTable(true)
	db.AutoMigrate(
		&models.Pegawai{},
		&models.Jabatan{},
		&models.User{},
		&models.Pendidikan{},
		&models.Prestasi{},
		&models.Tingkat{},
		&models.Pengumuman{},
		&models.Alamat{},
	)

}
