package db

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

type IDatabase interface {
	GetSGBD() gorm.Dialector
}

type MysqlDatabase struct {
	DbName     string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
}

func (m MysqlDatabase) GetSGBD() gorm.Dialector {

	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.DbUsername,
		m.DbPassword,
		m.DbHost,
		m.DbPort,
		m.DbName,
	))

}

func Connect(db IDatabase) {

	once.Do(func() {

		dbCon, err := gorm.Open(db.GetSGBD(), &gorm.Config{})

		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		instance = dbCon

	})

}

func GetInstance() *gorm.DB {
	return instance
}
