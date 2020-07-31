package fault

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorNotFound(t *testing.T) {
	msg := "test_error_msg"
	err := NewNotFound(msg)

	assert.NotEqual(t, nil, err)
	if err != nil {
		assert.Equal(t, msg, err.Error())
	}
	convErr, isOk := err.(*NotFond)
	assert.True(t, isOk, convErr)
}
