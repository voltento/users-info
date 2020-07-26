package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
	"net/http"
)

func (s *Service) GetUsers(ctx *gin.Context) {
	userParam := bindUser(ctx)
	users, err := s.storage.Users(userParam)
	if err != nil {
		err = errors.Wrap(err, "processing getUsers")
		s.logger.Error(err.Error())
		errorMsg := model.Error{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		ctx.JSON(http.StatusBadRequest, errorMsg)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
