package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)



type Datastorer interface {
	InitSchema()
}

type DB struct {
	*gorm.DB
}

func NewDB(dialect, dbName string) (*DB, error) {
	db, err := gorm.Open(dialect,os.Getenv("dbUser")+":"+ os.Getenv("dbPassword") + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) InitSchema() {
	db.AutoMigrate(&Favorite{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Tag{})

}