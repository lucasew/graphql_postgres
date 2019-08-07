package db

import (
	// import suport for postgress
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB initialized gorm database ready to use by the app
var DB *gorm.DB

// Ensure the migrations happen after DB initialization
var dbDone = make(chan (interface{}))

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=lucas59356 dbname=gorm sslmode=disable")
	DB.LogMode(true)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			dbDone <- nil
		}
	}()
}
