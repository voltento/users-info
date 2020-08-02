package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (suite *ServiceTestSuite) TestService_DeleteUserDeletedOk() {
	suite.userIdToUserData["1"] = &suite.testUser1
	path := suite.url + "/user/1"
	assert.Equal(suite.T(), http.StatusOK, httpDelete(path))
}

func (suite *ServiceTestSuite) TestService_DeleteUserNotDeletedOk() {
	path := suite.url + "/user/1"
	assert.Equal(suite.T(), http.StatusNoContent, httpDelete(path))
}

func (suite *ServiceTestSuite) TestService_DeleteUserBadRequest() {
	path := suite.url + "/user/a"
	assert.Equal(suite.T(), http.StatusBadRequest, httpDelete(path))
}
