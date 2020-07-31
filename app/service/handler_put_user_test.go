package service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
)

func (suite *ServiceTestSuite) TestService_PutUserOk() {
	path := suite.url + "/user/1?first_name=jj&last_name=vac&email=foo@mail&country_code=EN"
	assert.Equal(suite.T(), http.StatusOK, httpPut(path))
}

func (suite *ServiceTestSuite) TestService_PutUserNoFirstName() {
	path := suite.url + "/user/1?last_name=vac&email=foo@mail&country_code=EN"
	assert.Equal(suite.T(), http.StatusBadRequest, httpPut(path))
}

func (suite *ServiceTestSuite) TestService_PutUserNoLastName() {
	path := suite.url + "/user/1?last_name=vac&email=foo@mail&country_code=EN"
	assert.Equal(suite.T(), http.StatusBadRequest, httpPut(path))
}

func (suite *ServiceTestSuite) TestService_PutUserNoEmail() {
	path := suite.url + "/user/1?first_name=jj&last_name=vac@mail&country_code=EN"
	assert.Equal(suite.T(), http.StatusBadRequest, httpPut(path))
}

func (suite *ServiceTestSuite) TestService_PutUserNoCountryCode() {
	path := suite.url + "/user/1?first_name=jj&last_name=vac&email=foo@mail"
	assert.Equal(suite.T(), http.StatusBadRequest, httpPut(path))
}
