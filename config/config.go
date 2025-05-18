package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitSqlite()

	if err != nil {
		return fmt.Errorf("Error initializing SQLite database: %v", err)
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	return NewLogger(p)
}
