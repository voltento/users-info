package service

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/users-info/app/model"
)

func bindUser(ctx *gin.Context) *model.User {
	return &model.User{
		UserId:      ctx.DefaultQuery("user_id", ""),
		FirstName:   ctx.DefaultQuery("first_name", ""),
		SecondName:  ctx.DefaultQuery("second_name", ""),
		Email:       ctx.DefaultQuery("email", ""),
		CountryCode: ctx.DefaultQuery("country_code", ""),
	}
}
