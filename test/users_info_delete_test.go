package test_service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (s *UsersInfoTestSuite) TestUsersInfo_DeleteUserOk() {
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user1))
	_, users, err := s.GetUsers()
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	assert.Equal(s.T(), 1, len(users))
	assert.Equal(s.T(), http.StatusOK, s.DeleteUser(users[0].UserId))
	assert.Equal(s.T(), http.StatusNoContent, s.DeleteUser(users[0].UserId))
}

func (s *UsersInfoTestSuite) TestUsersInfo_DeleteUserBadId() {
	assert.Equal(s.T(), http.StatusBadRequest, s.DeleteUser("a"))
}
