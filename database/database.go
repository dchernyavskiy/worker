package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"worker/models"
)

var DB *gorm.DB

func Init() {
	dsn := "host=postgres.ewaybgadf8f6ared.eastus.azurecontainer.io user=postgres password=postgres port=5432"
	dsn2 := "host=postgres.ewaybgadf8f6ared.eastus.azurecontainer.io user=postgres password=postgres dbname=worker port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	_ = db.Exec("CREATE DATABASE worker")

	db, err = gorm.Open(postgres.Open(dsn2), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
	AutoMigrateModels()
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func AutoMigrateModels() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
}
