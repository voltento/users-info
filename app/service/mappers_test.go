package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/voltento/users-info/app/ui_errors"
	"net/http"
	"testing"
)

func Test_errToStatusCodeOk(t *testing.T) {
	assert.Equal(t, errToStatusCode(nil), http.StatusOK)
}

func Test_errToStatusCodeNoContent(t *testing.T) {
	assert.Equal(t, errToStatusCode(ui_errors.NewErrorNotFound("")), http.StatusNoContent)
}

func Test_errToStatusCodeBadRequest(t *testing.T) {
	assert.Equal(t, errToStatusCode(ui_errors.NewErrorBadRequest("")), http.StatusBadRequest)
}

func Test_errToStatusCodeBadRequestUnknownError(t *testing.T) {
	assert.Equal(t, errToStatusCode(errors.New("")), http.StatusInternalServerError)
}
