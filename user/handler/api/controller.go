package userAPI

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/stokbarang/config"
	"github.com/rachmankamil/stokbarang/user/domain"
)

type UserHandler struct {
	service domain.Service
}

func NewUserHandler(service domain.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (uh UserHandler) HealthCheck(ctx echo.Context) error {
	fmt.Println(config.Conf)
	return ctx.HTML(http.StatusOK, "OK HealthCheck + ")
}
