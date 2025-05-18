package config

import (
	"os"

	"github.com/HarlonGarcia/golang-jobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() (*gorm.DB, error) {
	logger = GetLogger("sqlite")
	logger.Debug("Initializing SQLite database")

	databasePath := "./db/main.db"

	_, err := os.Stat(databasePath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it")
		err := os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Error creating SQLite new database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(databasePath)

		if err != nil {
			logger.Errorf("Error creating SQLite new database file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Error automigrating SQLite database: %v", err)
		return nil, err
	}

	return db, nil
}
