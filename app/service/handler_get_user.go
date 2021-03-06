package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) GetUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	user, err := s.storage.User(id)
	if err != nil {
		s.logger.Named("get_user").Errorf("error: %v", err.Error())
		putErrToCtx(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
