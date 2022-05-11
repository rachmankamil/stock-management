package userAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/stokbarang/config"
	"github.com/rachmankamil/stokbarang/helper"
	"github.com/rachmankamil/stokbarang/user/domain"
)

type UserHandler struct {
	service    domain.Service
	validation *validator.Validate
}

func NewUserHandler(service domain.Service) UserHandler {

	return UserHandler{
		service:    service,
		validation: validator.New(),
	}
}

func (uh UserHandler) HealthCheck(ctx echo.Context) error {
	fmt.Println(config.Conf)
	return ctx.HTML(http.StatusOK, "OK HealthCheck + ")
}

func (uh UserHandler) AddUser(ctx echo.Context) error {
	req := RequestJSON{}
	ctx.Bind(&req)

	err := uh.validation.Struct(req)
	if err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	return ctx.HTML(http.StatusOK, "OK HealthCheck + ")
}

func (uh UserHandler) GetCost(ctx echo.Context) error {
	req := RequestOngkirJSON{}
	ctx.Bind(&req)

	err := uh.validation.Struct(req)
	if err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"message": stringerr})
	}

	param := fmt.Sprintf("origin=%d&destination=%d&weight=%d&courier=%s",
		req.Origin, req.Destination, req.Weight, req.Curier)

	httpreq, _ := http.NewRequest("POST", helper.RajaOngkirBaseURL+"cost", strings.NewReader(param))
	httpreq.Header.Add("key", "677971ea001e7e7d868fc03c52412452")
	httpreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpres, _ := http.DefaultClient.Do(httpreq)
	defer httpres.Body.Close()

	body, _ := ioutil.ReadAll(httpres.Body)

	respRO := ResponseRajaOngkir{}
	_ = json.Unmarshal(body, &respRO)

	return ctx.JSON(http.StatusOK, respRO.Rajaongkir.Results)
}
