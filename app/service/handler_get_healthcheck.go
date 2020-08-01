package service

import (
	hc "github.com/etherlabsio/healthcheck"
	"github.com/gin-gonic/gin"
)

func (s *Service) buildGetHealthCheck() func(*gin.Context) {
	opts := []hc.Option{
		hc.WithChecker(
			"database", s.storage,
		),
	}

	return gin.WrapH(hc.Handler(opts...))
}
