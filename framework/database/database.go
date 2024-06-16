package database

import (
	"log"

	"github.com/joaodeluchi/ms-users/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DbType        string
	Debug         bool
	AutomigrateDB bool
	Env           string
}

// Miss .Env file
func NewDb() *gorm.DB {
	dbInstance := &Database{}
	dbInstance.Env = "test"
	dbInstance.AutomigrateDB = true
	dbInstance.Debug = true
	dbInstance.Dsn = ":memory:"

	conn, err := dbInstance.Connect()

	if err != nil {
		log.Fatal("Database connection error")
	}

	return conn
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	d.Db, err = gorm.Open(sqlite.Open(d.Dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if d.AutomigrateDB {
		d.Db.AutoMigrate(&domain.User{})
	}

	return d.Db, nil

}
