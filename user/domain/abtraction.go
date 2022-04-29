package domain

type Service interface {
	//CreateToken this function for create token for jwt auth
	CreateToken(username, password string) (token string, err error)
	//InsertData this function is used when inserting new user
	InsertData(domain User) (response User, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain User) (id int, err error)
	Delete(id int) (err error)
	GetByUsernamePassword(username, password string) (domain User, err error)
	GetByID(id int) (domain User, err error)
}

type LocationRepo interface {
	GetLocation(lat, long float64) (city string)
}
