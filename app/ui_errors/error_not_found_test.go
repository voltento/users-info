package ui_errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorNotFound(t *testing.T) {
	msg := "test_error_msg"
	err := NewErrorNotFound(msg)

	assert.NotEqual(t, nil, err)
	if err != nil {
		assert.Equal(t, msg, err.Error())
	}
	convErr, isOk := err.(*ErrorNotFond)
	assert.True(t, isOk, convErr)
}
