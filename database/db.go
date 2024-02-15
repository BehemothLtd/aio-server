package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	if Db != nil {
		return Db
	}

	var err error
	newLogger := createLogger()

	Db, err = gorm.Open(mysql.Open(dbConnectString()), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return Db
}

func createLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Don't ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Include params in the SQL log
			Colorful:                  true,        // Enable color
		},
	)
}

func dbConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
}
