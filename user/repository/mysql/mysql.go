package repository

import (
	"github.com/rachmankamil/stokbarang/user/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (ur userRepo) Delete(id int) (err error) {
	return ur.DB.Delete("id = ?", id).Error
}

// GetByID implements domain.Repository
func (ur userRepo) GetByID(id int) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.Find("id = ?", id).First(&newRecord).Error

	return toDomain(newRecord), err
}

// GetByUsernamePassword implements domain.Repository
func (userRepo) GetByUsernamePassword(username string, password string) (domain domain.User, err error) {
	panic("unimplemented")
}

// Save implements domain.Repository
func (ur userRepo) Save(domain domain.User) (id int, err error) {
	record := fromDomain(domain)
	err = ur.DB.Create(&record).Error

	return int(record.ID), err
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepo{
		DB: db,
	}
}
