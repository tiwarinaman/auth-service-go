package database

import (
	"auth/model"
	"auth/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {

	cfg := config.Get()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
		return
	}

	DB = db

	migrateDatabase(DB)
}

func migrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error during database migration: %v", err)
	}
	log.Println("Database migration completed successfully.")
}
