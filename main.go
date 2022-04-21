package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/stokbarang/config"
	"github.com/rachmankamil/stokbarang/controller"
)

func main() {
	e := echo.New()
	config.Init()

	e.GET("/healthcheck", controller.HealthCheck)

	e.Start(":8080")
}
