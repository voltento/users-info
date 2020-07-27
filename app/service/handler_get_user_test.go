package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (suite *ServiceTestSuite) TestService_GetUserOk() {
	suite.userIdToUserData["1"] = &suite.testUser1

	query := suite.url + "/user/1"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUserFromHttpResponse(suite.testUser1, resp)
}

func (suite *ServiceTestSuite) TestService_GetUserNotFound() {
	query := suite.url + "/user/1"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusNoContent, resp.StatusCode)
}
