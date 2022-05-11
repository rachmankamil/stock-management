package domain

import "time"

type User struct {
	Username string
	Password string
	Fullname string
	DOB      time.Time
	Product  []Product
	Gender   string
}

type Product struct {
	Name  string
	Stock int
	Code  string
}
