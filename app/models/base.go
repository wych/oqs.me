package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"oqs.me/config"
)

var DB *gorm.DB

func DBInit() {
	db := config.Conf.DB
	var err error
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Addr, db.DatabaseName)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database initialize error: %s", err.Error())
	}
}

func GetMaxOQSID() (uint, error) {
	var record OQSRecord
	err := DB.Order("id desc").First(&record).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return record.ID, nil
}
