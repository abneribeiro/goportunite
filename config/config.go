package config

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	// Initialize SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing SQLite: %v", err)
	}
	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

func GetSQLiteDB() *gorm.DB {
	return db
}
