package database

import (
	"context"

	"gorm.io/gorm"
)

type DB interface {
	Create(interface{}) error
	Save(interface{}) error
	First(interface{}, ...interface{}) error
	Delete(interface{}, ...interface{}) error
	Where(interface{}, ...interface{}) DB
	Order(value interface{}) DB
	Model(value interface{}) DB
	Count(count *int64) error
	WithContext(c context.Context) DB
}

type db struct {
	*gorm.DB
}

func NewDB(gormDB *gorm.DB) *db {
	return &db{DB: gormDB}
}

func (s *db) Create(any interface{}) error {
	return s.DB.Create(any).Error
}

func (s *db) Save(any interface{}) error {
	return s.DB.Save(any).Error
}

func (s *db) First(out interface{}, where ...interface{}) error {
	return s.DB.First(out, where...).Error
}

func (s *db) Delete(any interface{}, where ...interface{}) error {
	return s.DB.Delete(any, where...).Error
}

func (s *db) Where(query interface{}, args ...interface{}) DB {
	return &db{DB: s.DB.Where(query, args...)}
}

func (s *db) Order(value interface{}) DB {
	return &db{DB: s.DB.Order(value)}
}

func (s *db) Model(value interface{}) DB {
	return &db{DB: s.DB.Model(value)}
}

func (s *db) Count(count *int64) error { // Adjusted to return an error directly
	return s.DB.Count(count).Error
}

func (s *db) WithContext(c context.Context) DB {
	return &db{DB: s.DB.WithContext(c)}
}
