package database

import "gorm.io/gorm"

type Client interface {
	Connect()
	Migrate()
	GetDb() *gorm.DB
}
