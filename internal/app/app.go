package app

import (
	"context"
	"database/sql"

	"github.com/core-go/health/gin"
	s "github.com/core-go/health/sql"
	"github.com/core-go/log/zap"

	"go-service/internal/user"
)

type ApplicationContext struct {
	Health *gin.Handler
	User   user.UserTransport
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	db, err := sql.Open(cfg.Sql.Driver, cfg.Sql.DataSourceName)
	if err != nil {
		return nil, err
	}
	logError := log.LogError

	userHandler, err := user.NewUserHandler(db, logError)
	if err != nil {
		return nil, err
	}

	sqlChecker := s.NewHealthChecker(db)
	healthHandler := gin.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
