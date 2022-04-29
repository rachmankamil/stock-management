package repository

import (
	"time"

	"github.com/rachmankamil/stokbarang/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	DOB      time.Time
}

func toDomain(rec User) domain.User {
	return domain.User{
		Username: rec.Username,
		Password: rec.Password,
		Fullname: rec.Fullname,
		DOB:      rec.DOB,
	}
}

func fromDomain(rec domain.User) User {
	return User{
		Username: rec.Username,
		Password: rec.Password,
		Fullname: rec.Fullname,
		DOB:      rec.DOB,
	}
}
