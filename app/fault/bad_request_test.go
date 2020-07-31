package fault

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorBadRequest(t *testing.T) {
	msg := "test_error_msg"
	err := NewBadRequest(msg)

	assert.NotEqual(t, nil, err)
	if err != nil {
		assert.Equal(t, msg, err.Error())
	}
	convErr, isOk := err.(*BadRequest)
	assert.True(t, isOk, convErr)
}
