package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	user, err := s.storage.User(id)
	if err != nil {
		putErrToCtx(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
