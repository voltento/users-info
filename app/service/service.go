package service

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/users-info/app/connectors"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	storage connectors.Storage
	config  *Config
	engine  *gin.Engine
	logger  *zap.SugaredLogger
}

func NewService(config *Config, logger *zap.SugaredLogger, storage connectors.Storage) (error, *Service) {
	var engine *gin.Engine
	if !config.LogGinGonic {
		engine = gin.New()
		engine.Use(gin.Recovery())

		// Redirect logging from gin to zap
		engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			logMethod := logger.Infow
			if len(param.ErrorMessage) > 0 {
				logMethod = logger.Errorw
			}

			logMethod("",
				"client_ip", param.ClientIP,
				"timestamp", param.TimeStamp.Format(time.RFC1123),
				"method", param.Method,
				"path", param.Path,
				"proto", param.Request.Proto,
				"status_code", param.StatusCode,
				"latency", param.Latency,
				"error", param.ErrorMessage,
			)
			return ""
		}))
	} else {
		engine = gin.Default()
	}

	s := &Service{
		storage: storage,
		config:  config,
		engine:  engine,
		logger:  logger,
	}

	s.ConnectHandlers()

	return nil, s
}

func (s *Service) Run() error {
	er := s.engine.Run(s.config.Address)
	s.Stop()
	return er
}

func (s *Service) ConnectHandlers() {
	s.engine.GET("/users", s.GetUsers)
	s.engine.GET("/user/:user_id", s.GetUser)
}

func (s *Service) Stop() {
	er := s.storage.Stop()
	if er != nil {
		s.logger.Warn("error occured during stop server", "error", er.Error())
	}
}
