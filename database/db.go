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

func dbConnectString() string {
	var DB_USERNAME = os.Getenv("DB_USERNAME")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_DATABASE = os.Getenv("DB_DATABASE")
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")

	return DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?" + "parseTime=true&loc=Local"
}

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Dont Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dbConnectString()), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)

		return nil
	}

	return db
}
