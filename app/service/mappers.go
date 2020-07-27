package service

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/users-info/app/model"
	"github.com/voltento/users-info/app/ui_errors"
	"net/http"
)

func ctxToUser(ctx *gin.Context) *model.User {
	return &model.User{
		UserId:      ctx.DefaultQuery("user_id", ""),
		FirstName:   ctx.DefaultQuery("first_name", ""),
		SecondName:  ctx.DefaultQuery("second_name", ""),
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
	case *ui_errors.ErrorNotFond:
		return http.StatusNoContent
	case *ui_errors.ErrorBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
