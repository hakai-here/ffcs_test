package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	db = Connect()
	return db
}

func Connect() *gorm.DB {
	username := viper.GetString("POSTGRES_USER")
	password := viper.GetString("POSTGRES_PASSWORD")
	database := viper.GetString("POSTGRES_DB")
	host := viper.GetString("POSTGRES_HOST")
	// creating the uri
	dburi := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=disable password=%s", host, username, database, password)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dburi,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("[DB] Error : %s", err.Error())
	}
	return db

}
