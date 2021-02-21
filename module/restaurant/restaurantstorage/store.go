package restaurantstorage

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

// NewSQLStore create new store
func NewSQLStore(db *gorm.DB) *store {
	return &store{db: db}
}
