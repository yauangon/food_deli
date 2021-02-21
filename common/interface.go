package common

import "gorm.io/gorm"

type DBProvider interface {
	GetMainDBConnection() *gorm.DB
}
