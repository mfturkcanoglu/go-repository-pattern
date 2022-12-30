package database

import (
	"log"

	"github.com/mfturkcanoglu/go-repository-pattern/pkg/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteClient struct {
	logger *log.Logger
	Db     *gorm.DB
}

func NewSqliteClient(l *log.Logger) Client {
	return &SqliteClient{
		logger: l,
		Db:     nil,
	}
}

func (c *SqliteClient) Connect() {
	db, err := gorm.Open(sqlite.Open("../db.db"), &gorm.Config{})
	if err != nil {
		c.logger.Fatalln("Failed to connect db")
	}
	c.Db = db
	c.logger.Println("Connected to db")
}

func (c *SqliteClient) Migrate() {
	c.Db.AutoMigrate(entity.Todo{})
}

func (c *SqliteClient) GetDb() *gorm.DB {
	return c.Db
}
