package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
	"net/http"
)

func (s *Service) GetUser(ctx *gin.Context) {
	userParam := bindUser(ctx)

	user, err := s.storage.User(userParam)
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

	ctx.JSON(http.StatusOK, user)
}
