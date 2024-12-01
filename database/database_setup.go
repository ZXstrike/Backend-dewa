package database

import (
	"zxsttm/database/models"
	"zxsttm/server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MySQLConnect(config *config.MySQLConfig) (*gorm.DB, error) {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migration(db)

	DB = db

	return db, nil
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.ESP{},
		&models.DHT{},
		&models.PZEM{},
		&models.Relay{},
		&models.UserOTP{},
	)
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("Database connection is not initialized")
	}
	return DB
}