package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/stokbarang/config"
)

func HealthCheck(ctx echo.Context) error {
	fmt.Println(config.Conf)
	return ctx.HTML(http.StatusOK, "OK HealthCheck + ")
}
