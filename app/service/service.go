package service

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/users-info/app/connectors"
)

type Service struct {
	storage connectors.Storage
	config  Config
	engine  *gin.Engine
}

func NewService(config Config) (error, *Service) {
	var engine *gin.Engine
	if !config.LogGinGonic {
		engine = gin.New()
		engine.Use(gin.Recovery())
	} else {
		engine = gin.Default()
	}

	s := &Service{
		storage: config.Storage,
		config:  config,
		engine:  engine,
	}

	s.ConnectHandlers()

	return nil, s
}

func (s *Service) Run() error {
	return s.engine.Run(s.config.Addr)
}

func (s *Service) ConnectHandlers() {
	s.engine.GET("/users", s.GetUsers)
}
