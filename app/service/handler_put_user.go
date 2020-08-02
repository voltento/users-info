package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) PutUser(ctx *gin.Context) {
	userParam := ctxPathToUser(ctx)
	userParam.UserId = ctx.Param("user_id")

	err := s.storage.UpdateUser(userParam)
	if err != nil {
		putErrToCtx(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
