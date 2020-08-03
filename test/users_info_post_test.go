package test_service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (s *UsersInfoTestSuite) TestUsersInfo_PostUser() {
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user1))
	assert.Equal(s.T(), http.StatusBadRequest, s.PostUser(s.user1))
}
