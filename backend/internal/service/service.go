package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mishaRomanov/maglev/internal/config"
	"github.com/mishaRomanov/maglev/internal/handlers"
	"github.com/mishaRomanov/maglev/internal/middleware"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	app      *gin.Engine
	cfg      *config.GlobalConfig
	handlers *handlers.Handler
}

func New(cfg *config.GlobalConfig,
	redisClient *redis.Client,
) *Service {
	backend := Service{
		gin.New(),
		cfg,
		handlers.New(redisClient),
	}
	return &backend
}

func (s *Service) Run() error {
	s.app.Use(middleware.RequestLog())

	s.app.GET("/", s.handlers.MainPage)
	s.app.POST("/register", s.handlers.Register)

	if runErr := s.app.Run(":5353"); runErr != nil {
		return runErr
	}
	return nil
}
