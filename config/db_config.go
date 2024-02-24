package config

import (
	"github.com/ujjwal8007/models/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgreSQL() (*gorm.DB, error) {
	dsn := "host=" + "127.0.0.1" +
		" user=" + "postgres" +
		" password=" + "postgres" +
		" dbname=" + "dblrucache" +
		" port=" + "5432" +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.LRUCache{})
	return db, err
}
