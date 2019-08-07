package db

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
)

var DB *gorm.DB
// Sinaliza que o db foi inicializado, garante que as AutoMigrations não vão rodar até o banco estar inicializado
var dbDone = make(chan(interface{}))

func init() {
    var err error
    DB, err = gorm.Open("postgres", "host=localhost user=lucas59356 dbname=gorm sslmode=disable")
    DB.LogMode(true)
    if err != nil {
        panic(err)
    }
    go func () {
        for {
            dbDone <- nil
        }
    }()
}
