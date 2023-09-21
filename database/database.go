package database

import (
	"log"
	"os"
	"sync"

	"github.com/muchrief/go_pijar/database/db_model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dbOnce sync.Once

type DatabaseType string

const (
	MY_SQL     DatabaseType = "my_sql"
	POSTGRESQL DatabaseType = "postgres"
	SQLITE     DatabaseType = "sqlite"
)

func InitializeDB() {
	dbOnce.Do(func() {
		uri := getDatabaseUri()
		db, err := openConnectionDatabase(POSTGRESQL, uri, &gorm.Config{})
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		err = autoMigrate(db)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		DB = db
	})
}

func getDatabaseUri() string {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		return "host=localhost user=postgres password=m03kht4r1999 dbname=pijar port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	}

	return uri
}

func openConnectionDatabase(dbType DatabaseType, uri string, config *gorm.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch dbType {
	case MY_SQL:
		db, err = gorm.Open(mysql.Open(uri), config)
	case POSTGRESQL:
		db, err = gorm.Open(postgres.Open(uri), config)
	case SQLITE:
		db, err = gorm.Open(sqlite.Open(uri), config)
	}

	return db, err
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&db_model.User{})
	if err != nil {
		return err
	}

	return nil
}

// func autoDowngrade(db *gorm.DB) error {
// 	err := db.Migrator().DropTable(&db_model.User{})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
