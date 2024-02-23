package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ujjwal8007/models"
	"github.com/ujjwal8007/models/entity"
	"github.com/ujjwal8007/store"
)

type LRUCacheService interface {
	Put(ctx context.Context, req models.SetKeyRequest) error
	Get(ctx context.Context, key string) (models.GetKeyResponse, error)
}

type LRUCache struct {
	lruCacheStore store.LRUCacheStore
	capacity      int
}

func NewLRUCache(lruCacheStore store.LRUCacheStore, c int) *LRUCache {
	return &LRUCache{
		lruCacheStore: lruCacheStore,
		capacity:      c,
	}
}

func (lru *LRUCache) Put(ctx context.Context, req models.SetKeyRequest) error {
	// Check if the key already exists in the cache
	entityData, err := lru.lruCacheStore.GetEntityByKey(ctx, req.Key)
	if err == nil {
		// If the key exists, update the value and expiry time
		entityData.Value = req.Value
		entityData.Expiry = time.Now().Add(time.Duration(req.ExpiryTime) * time.Second)
		if err := lru.lruCacheStore.SaveEntity(ctx, entityData); err != nil {
			return fmt.Errorf("failed to SaveEntity with err %w", err)
		}
	} else {
		// If the key does not exist, create a new entry
		newEntity := entity.LRUCache{
			Key:      req.Key,
			Value:    req.Value,
			LastUsed: time.Now(),
			Expiry:   time.Now().Add(time.Duration(req.ExpiryTime) * time.Second),
		}
		if err := lru.lruCacheStore.CreateEntity(ctx, newEntity); err != nil {
			return fmt.Errorf("failed to CreateEntity with err %w", err)
		}
	}

	// Check if the cache has reached its capacity
	count, err := lru.lruCacheStore.CountEntities(ctx)
	if err == nil && count > int64(lru.capacity) {
		// Evict the least recently used items
		leastRecentlyUsedEntity, err := lru.lruCacheStore.GetLeastRecentlyUsedEntity(ctx)
		if err == nil {
			if err := lru.lruCacheStore.DeleteEntity(ctx, leastRecentlyUsedEntity); err != nil {
				return fmt.Errorf("failed to DeleteEntity with err %w", err)
			}
		}
	}
	return nil
}

func (lru *LRUCache) Get(ctx context.Context, key string) (models.GetKeyResponse, error) {
	lru.lruCacheStore.DeleteExpiredKeys(ctx)
	entity, err := lru.lruCacheStore.GetEntityByKey(ctx, key)
	if err != nil {
		// Key not found
		return models.GetKeyResponse{}, fmt.Errorf("failed to GetEntityByKey %w", err)
	}
	// Update the LastUsed timestamp
	entity.LastUsed = time.Now()
	if err := lru.lruCacheStore.SaveEntity(ctx, entity); err != nil {
		return models.GetKeyResponse{}, fmt.Errorf("failed to SaveEntity %w", err)
	}

	return models.GetKeyResponse{
		Key:      entity.Key,
		Value:    entity.Value,
		LastUsed: entity.LastUsed.Format("02/01/2006   15:04:05"),
	}, nil
}
