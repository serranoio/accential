package database

import (
	"log"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() {
	dir, err := os.Getwd()

	if err != nil {
		log.Panicln(err)
	}

	Db, err = gorm.Open(sqlite.Open(path.Join(dir, "database", "database.db")), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

}
