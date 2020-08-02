package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) GetUsers(ctx *gin.Context) {
	userParam := ctxPathToUser(ctx)
	users, err := s.storage.Users(userParam)
	if err != nil {
		s.logger.Named("get_users").Errorf("error: %v", err.Error())
		putErrToCtx(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
