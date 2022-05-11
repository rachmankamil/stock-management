package userAPI

import (
	"time"

	"github.com/rachmankamil/stokbarang/user/domain"
)

type RequestJSON struct {
	Username string    `json:"username" validate:"required,email"`
	Password string    `json:"password" validate:"required"`
	DOB      time.Time `json:"day_of_birth"`
	Fullname string    `json:"fullname"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Username: req.Username,
	}
}

type ResponseJSON struct {
	Username string    `json:"username"`
	DOB      time.Time `json:"day_of_birth"`
	Fullname string    `json:"fullname"`
}

func fromDomain(domain domain.User) ResponseJSON {
	return ResponseJSON{
		Username: domain.Username,
	}
}

type RequestOngkirJSON struct {
	Origin      int    `json:"origin" validate:"required,number"`
	Destination int    `json:"dest" validate:"required,number"`
	Weight      int    `json:"weight" validate:"required,number"`
	Curier      string `json:"curier" validate:"required,oneof=jne pos tiki"`
}

type ResponseRajaOngkir struct {
	Rajaongkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"destination_details"`
		Results []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}
