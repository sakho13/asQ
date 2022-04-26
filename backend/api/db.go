package api

import (
	"errors"
	"log"
	"time"

	"github.com/sakho13/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBInit executes initializing DB.
func DBInit() {
	localDsn := "root:MSAKHO13P@tcp(db_mysql)/with_me_db?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := dbOpen(localDsn, 30)
	if err != nil {
		panic(err)
	}

	migrations(db)

	DB = db

}
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
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
