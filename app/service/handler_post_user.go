package service

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/users-info/app/model"
	"net/http"
)

func (s *Service) PostUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		putErrToCtx(err, ctx)
		return
	}

	err := s.storage.AddUser(&user)
	if err != nil {
		putErrToCtx(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
