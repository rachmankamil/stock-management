package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/stokbarang/config"

	"github.com/rachmankamil/stokbarang/user"
)

func main() {

	config.Init()
	db := config.DBInit()

	user := user.NewUserFactory(db)

	//Route
	e := echo.New()

	e.GET("/healthcheck", user.HealthCheck)
	e.POST("/user", user.AddUser)
	e.POST("/ro", user.GetCost)

	e.Start(":8080")
}
