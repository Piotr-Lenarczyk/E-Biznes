package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "github.com/labstack/echo/v4"
    "go-echo-gorm-app/controllers" // import controllers
)

var db *gorm.DB
var err error

func init() {
    // Setup SQLite database
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Error opening database:", err)
    }

    // Migrate the schema
    db.AutoMigrate(&controllers.Product{})

    // Set DB in the controller
    controllers.SetDB(db)
}

func main() {
    e := echo.New()

    // Register the routes
    controllers.RegisterProductRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
