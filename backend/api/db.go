package api

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sakho13/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBInit executes initializing DB.
func DBInit() {
	dsn := os.Getenv("PLANETSCALE_DSN")
	if len(strings.TrimSpace(dsn)) == 0 {
		panic("DSN未設定")
	}

	log.Printf("Connecting to %v.\n", dsn)

	db, err := dbOpen(dsn, 30)
	if err != nil {
		panic(err)
	}

	migrations(db)

	DB = db

}
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		// &models.Post{},
		// &models.Diary{},
		// &models.Genre{},

		// For Admin
		// &models.Admin{},
		// &models.Letter{},
	)
	if err != nil {
		panic(err)
	}
}

func dbOpen(path string, tryCount uint) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil && db != nil {
		if tryCount == 0 {
			return nil, errors.New("TryCount over")
		}

		log.Printf("Making the connection. (%v)\n", tryCount)
		time.Sleep(time.Second * 5)

		tryCount--
		return dbOpen(path, tryCount)
	}

	log.Printf("The connection was made.\n")
	return db, nil
}
