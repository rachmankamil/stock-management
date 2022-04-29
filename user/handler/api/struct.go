package userAPI

import (
	"time"

	"github.com/rachmankamil/stokbarang/user/domain"
)

type RequestJSON struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
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
