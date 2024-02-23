package entity

import (
	"time"
)

type LRUCache struct {
	ID       int       `gorm:"column:lru_cache_id;primary_key;AUTO_INCREMENT"`
	Key      string    `gorm:"column:lru_key;unique;not null;type:varchar(64)"`
	Value    string    `gorm:"column:value;not null;type:varchar(64)"`
	LastUsed time.Time `gorm:"column:last_used"`
	Expiry   time.Time `gorm:"column:expiry"`
}

func (m *LRUCache) TableName() string {
	return "lru_cache"
}
