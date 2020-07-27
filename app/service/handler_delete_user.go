package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	err := s.storage.DropUser(id)
	if err != nil {
		putErrToCtx(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
