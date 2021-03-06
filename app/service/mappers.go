package service

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/users-info/app/fault"
	"github.com/voltento/users-info/app/model"
	"net/http"
)

func ctxPathToUser(ctx *gin.Context) *model.User {
	return &model.User{
		UserId:      ctx.DefaultQuery("user_id", ""),
		FirstName:   ctx.DefaultQuery("first_name", ""),
		LastName:    ctx.DefaultQuery("last_name", ""),
		Email:       ctx.DefaultQuery("email", ""),
		CountryCode: ctx.DefaultQuery("country_code", ""),
	}
}

func putErrToCtx(err error, ctx *gin.Context) {
	code := errToStatusCode(err)
	errorMsg := model.Error{
		Message: err.Error(),
		Code:    code,
	}
	ctx.JSON(code, errorMsg)
}

func errToStatusCode(err error) int {
	switch err.(type) {
	case nil:
		return http.StatusOK
	case *fault.NotFond:
		return http.StatusNoContent
	case *fault.BadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
