package api

import (
	"errors"
	"log"
	"time"

	"github.com/sakho13/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBInit executes initializing DB.
func DBInit() {
	// mysql
	// localDsn := "root:MSAKHO13P@tcp(db_mysql)/with_me_db?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	// postgres
	localDsn := "host=postgres port=5432 user=user password=password dbname=with_me_db sslmode=disable TimeZone=Asia/Tokyo"
	db, err := dbOpen(localDsn, 30)
	if err != nil {
		panic(err)
	}

	migrations(db)

	DB = db

}
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Diary{},
		&models.Genre{},

		// For Admin
		&models.Admin{},
		&models.Letter{},
	)
	if err != nil {
		panic(err)
	}
}

func dbOpen(path string, tryCount uint) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})
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
