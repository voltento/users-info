package ui_errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorBadRequest(t *testing.T) {
	msg := "test_error_msg"
	err := NewErrorBadRequest(msg)

	assert.NotEqual(t, nil, err)
	if err != nil {
		assert.Equal(t, msg, err.Error())
	}
	convErr, isOk := err.(*ErrorBadRequest)
	assert.True(t, isOk, convErr)
}
