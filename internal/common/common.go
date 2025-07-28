package common

import "gorm.io/gorm"

type DBInterface interface {
	ConnectDB() *gorm.DB
}
