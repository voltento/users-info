package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/voltento/users-info/app/fault"
	"net/http"
	"testing"
)

func Test_errToStatusCodeOk(t *testing.T) {
	assert.Equal(t, errToStatusCode(nil), http.StatusOK)
}

func Test_errToStatusCodeNoContent(t *testing.T) {
	assert.Equal(t, errToStatusCode(fault.NewNotFound("")), http.StatusNoContent)
}

func Test_errToStatusCodeBadRequest(t *testing.T) {
	assert.Equal(t, errToStatusCode(fault.NewBadRequest("")), http.StatusBadRequest)
}

func Test_errToStatusCodeBadRequestUnknownError(t *testing.T) {
	assert.Equal(t, errToStatusCode(errors.New("")), http.StatusInternalServerError)
}
