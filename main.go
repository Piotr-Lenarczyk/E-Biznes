package main

import (
	"go-echo-gorm-app/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	controllers.RegisterProductRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
