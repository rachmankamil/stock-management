package user

import (
	"gorm.io/gorm"

	userAPI "github.com/rachmankamil/stokbarang/user/handler/api"
	userRepoMySQL "github.com/rachmankamil/stokbarang/user/repository/mysql"
	userService "github.com/rachmankamil/stokbarang/user/service"
)

func NewUserFactory(db *gorm.DB) (userHandler userAPI.UserHandler) {
	userRepo := userRepoMySQL.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo)
	userHandler = userAPI.NewUserHandler(userServ)
	return
}
