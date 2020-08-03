package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) PutUser(ctx *gin.Context) {
	user := ctxPathToUser(ctx)
	user.UserId = ctx.Param("user_id")

	err := s.storage.UpdateUser(user)
	if err != nil {
		putErrToCtx(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
