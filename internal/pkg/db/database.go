package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDB() error {

	log.Println("Attempting DB connection...")

	//DatabaseUrl := "postgres://postgres:postgres@database:5432?sslmode=disable"
	DatabaseUrlDebug := "postgres://postgres:postgres@127.0.0.1:5432?sslmode=disable"

	sqldb, err := gorm.Open(postgres.Open(DatabaseUrlDebug), &gorm.Config{})
	if err != nil {
		return err
	}

	testDB, err := sqldb.DB()
	err = testDB.Ping()
	if err != nil {
		return err
	}

	Db = sqldb

	log.Println("Connected to DB....")
	return nil
}
