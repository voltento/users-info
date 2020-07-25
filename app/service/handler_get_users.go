package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) GetUsers(ctx *gin.Context) {
	err, users := s.storage.Users()
	if err != nil {
		// TODO
	}

	ctx.JSON(http.StatusOK, users)
}
