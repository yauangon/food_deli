package userstorage

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

// NewSQLStore for user
func NewSQLStore(db *gorm.DB) *store {
	return &store{db: db}
}
