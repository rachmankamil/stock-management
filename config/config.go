package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string

	JWTSecret string
}

var Conf Config

func Init() {
	Conf = Config{
		DBNAME:    os.Getenv("DBNAME"),
		DBUSER:    os.Getenv("DBUSER"),
		DBPASS:    os.Getenv("DBPASS"),
		DBHOST:    os.Getenv("DBHOST"),
		DBPORT:    os.Getenv("DBPORT"),
		JWTSecret: os.Getenv("JWTSECRET"),
	}
	fmt.Printf("%+v", Conf)
}

func DBInit() (DB *gorm.DB) {
	DB, _ = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
				Conf.DBUSER,
				Conf.DBPASS,
				Conf.DBHOST,
				Conf.DBPORT,
				Conf.DBNAME,
			),
		),
	)
	return
}
