package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func (s *Service) GetUsers(ctx *gin.Context) {
	users, err := s.storage.Users()
	if err != nil {
		s.logger.Error(errors.Wrap(err, "processing getUsers").Error())
	}

	ctx.JSON(http.StatusOK, users)
}
