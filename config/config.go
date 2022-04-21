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
	DBPORT int

	JWTSecret string
}

var Conf Config
var DB *gorm.DB

func Init() {
	Conf = Config{
		DBNAME:    os.Getenv("DBNAME"),
		DBUSER:    os.Getenv("DBUSER"),
		JWTSecret: os.Getenv("JWTSECRET"),
	}
	DBInit()
}

func DBInit() {
	DB, _ = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
				Conf.DBUSER,
				Conf.DBPASS,
				Conf.DBHOST,
				Conf.DBPORT,
				Conf.DBNAME,
			),
		),
	)
}
