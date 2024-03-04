// models/setup.go
package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=postgres dbname=go_transit port=5432 sslmode=disable timezone=America/Los_Angeles"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})  // change the database provider if necessary

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&Station{})  // register Post model

    DB = database
}