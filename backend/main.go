package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "go-echo-gorm-app/controllers"
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
    if err := db.AutoMigrate(&controllers.Product{}, &controllers.Cart{}, &controllers.Category{}, &controllers.Payment{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Set DB in the controller
    SetDB(db)
}

func SetDB(database *gorm.DB) {
    db = database
}

func main() {
    e := echo.New()

    e.Use(middleware.CORS())

    // Register the routes
    controllers.RegisterProductRoutes(e, db)
    controllers.RegisterCartRoutes(e, db)
    controllers.RegisterCategoryRoutes(e, db)
    controllers.RegisterPaymentRoutes(e, db)

    e.Logger.Fatal(e.Start(":8080"))
}
