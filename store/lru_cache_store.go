package store

import (
	"context"
	"time"

	"github.com/ujjwal8007/database"
	"github.com/ujjwal8007/models/entity"
)

type LRUCacheStore interface {
	GetEntityByKey(ctx context.Context, key string) (entity.LRUCache, error)
	SaveEntity(ctx context.Context, entity entity.LRUCache) error
	CreateEntity(ctx context.Context, entity entity.LRUCache) error
	CountEntities(ctx context.Context) (int64, error)
	GetLeastRecentlyUsedEntity(ctx context.Context) (entity.LRUCache, error)
	DeleteEntity(ctx context.Context, entity entity.LRUCache) error
	DeleteExpiredKeys(ctx context.Context) error
}

type lruCacheStore struct {
	dbStore database.DB
}

func NewLRUCacheStore(dbStore database.DB) LRUCacheStore {
	return &lruCacheStore{dbStore: dbStore}
}

func (s *lruCacheStore) GetEntityByKey(ctx context.Context, key string) (entity.LRUCache, error) {
	var entity entity.LRUCache
	if err := s.dbStore.WithContext(ctx).First(&entity, "lru_key = ?", key); err != nil {
		return entity, err // Corrected: return the zero value of entity.LRUCache directly
	}
	return entity, nil
}

func (s *lruCacheStore) SaveEntity(ctx context.Context, entity entity.LRUCache) error {
	if err := s.dbStore.WithContext(ctx).Save(&entity); err != nil {
		return err
	}
	return nil
}

func (s *lruCacheStore) CreateEntity(ctx context.Context, entity entity.LRUCache) error {
	if err := s.dbStore.WithContext(ctx).Create(&entity); err != nil {
		return err
	}
	return nil
}

func (s *lruCacheStore) CountEntities(ctx context.Context) (int64, error) {
	var count int64
	err := s.dbStore.WithContext(ctx).Model(&entity.LRUCache{}).Count(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *lruCacheStore) GetLeastRecentlyUsedEntity(ctx context.Context) (entity.LRUCache, error) {
	var entity entity.LRUCache
	if err := s.dbStore.WithContext(ctx).Order("last_used ASC").First(&entity); err != nil {
		return entity, err
	}
	return entity, nil
}

func (s *lruCacheStore) DeleteEntity(ctx context.Context, entity entity.LRUCache) error {
	if err := s.dbStore.WithContext(ctx).Delete(&entity); err != nil {
		return err
	}
	return nil
}
func (s *lruCacheStore) DeleteExpiredKeys(ctx context.Context) error {
	if err := s.dbStore.WithContext(ctx).Where("expiry < ?", time.Now()).Delete(&entity.LRUCache{}); err != nil {
		return err
	}
	return nil
}
