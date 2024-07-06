package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"go-service/internal/user/handler"
	"go-service/internal/user/repository/adapter"
	"go-service/internal/user/service"
)

type UserTransport interface {
	All(*gin.Context)
	Load(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Patch(*gin.Context)
	Delete(*gin.Context)
}

func NewUserHandler(db *sql.DB) (UserTransport, error) {
	userRepository, err := adapter.NewUserAdapter(db)
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return userHandler, nil
}
