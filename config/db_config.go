package config

import (
	"github.com/ujjwal8007/models/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgreSQL() (*gorm.DB, error) {
	dsn := "host=" + "lru-cache-instance.cxuuawwem0df.eu-north-1.rds.amazonaws.com" +
		" user=" + "postgres" +
		" password=" + "S21fe35218" +
		" dbname=" + "lru_cache_initial" +
		" port=" + "5432" +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.LRUCache{})
	return db, err
}
