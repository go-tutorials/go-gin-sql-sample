package user

import (
	"context"
	"database/sql"

	v "github.com/core-go/core/v10"
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

func NewUserHandler(db *sql.DB, logError func(context.Context, string, ...map[string]interface{})) (UserTransport, error) {
	validator, err := v.NewValidator()
	if err != nil {
		return nil, err
	}

	userRepository, err := adapter.NewUserAdapter(db)
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, logError, validator.Validate)
	return userHandler, nil
}
