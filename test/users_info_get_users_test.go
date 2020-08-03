package test_service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (s *UsersInfoTestSuite) TestUsersInfo_GetUsersNoUsers() {
	code, users, err := s.GetUsers()
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	assert.Equal(s.T(), http.StatusOK, code)
	assert.Equal(s.T(), 0, len(users))
}

func (s *UsersInfoTestSuite) TestUsersInfo_GetUsersOneUser() {
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user1))
	code, users, err := s.GetUsers()
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	assert.Equal(s.T(), http.StatusOK, code)
	assert.Equal(s.T(), 1, len(users))

	s.user1.UserId = users[0].UserId
	assert.Equal(s.T(), s.user1, users[0])
}

func (s *UsersInfoTestSuite) TestUsersInfo_GetUsersFirstName() {
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user1))
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user2))
	code, users, err := s.GetUsers("first_name=" + s.user1.FirstName)
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	assert.Equal(s.T(), http.StatusOK, code)
	assert.Equal(s.T(), 2, len(users))
}

func (s *UsersInfoTestSuite) TestUsersInfo_GetUsersLastName() {
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user1))
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user2))
	code, users, err := s.GetUsers("last_name=" + s.user1.LastName)
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	assert.Equal(s.T(), http.StatusOK, code)
	assert.Equal(s.T(), 1, len(users))

	s.user1.UserId = users[0].UserId
	assert.Equal(s.T(), s.user1, users[0])
}

func (s *UsersInfoTestSuite) TestUsersInfo_GetUsersCountryCode() {
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user1))
	assert.Equal(s.T(), http.StatusOK, s.PostUser(s.user2))
	code, users, err := s.GetUsers("country_code=" + s.user1.CountryCode)
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	assert.Equal(s.T(), http.StatusOK, code)
	assert.Equal(s.T(), 1, len(users))

	s.user1.UserId = users[0].UserId
	assert.Equal(s.T(), s.user1, users[0])
}
