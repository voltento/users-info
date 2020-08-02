package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/voltento/users-info/app/model"
	"net/http"
)

func (suite *ServiceTestSuite) TestService_PostUserOk() {
	path := suite.url + "/user/"
	user := &model.User{
		FirstName:   "Man",
		LastName:    "Lan",
		Email:       "foo@mail",
		CountryCode: "EN",
	}
	assert.Equal(suite.T(), http.StatusOK, httpPost(path, user))
}

func (suite *ServiceTestSuite) TestService_PostUserNoFirstName() {
	path := suite.url + "/user/"
	user := &model.User{
		LastName:    "Lan",
		Email:       "foo@mail",
		CountryCode: "EN",
	}
	assert.Equal(suite.T(), http.StatusBadRequest, httpPost(path, user))
}

func (suite *ServiceTestSuite) TestService_PostUserNoLastName() {
	path := suite.url + "/user/"
	user := &model.User{
		FirstName:   "Man",
		Email:       "foo@mail",
		CountryCode: "EN",
	}
	assert.Equal(suite.T(), http.StatusBadRequest, httpPost(path, user))
}

func (suite *ServiceTestSuite) TestService_PostUserNoEmail() {
	path := suite.url + "/user/"
	user := &model.User{
		FirstName:   "Man",
		LastName:    "Lan",
		CountryCode: "EN",
	}
	assert.Equal(suite.T(), http.StatusBadRequest, httpPost(path, user))
}

func (suite *ServiceTestSuite) TestService_PostUserNoCountryCode() {
	path := suite.url + "/user/"
	user := &model.User{
		FirstName: "Man",
		LastName:  "Lan",
		Email:     "foo@mail",
	}
	assert.Equal(suite.T(), http.StatusBadRequest, httpPost(path, user))
}

func (suite *ServiceTestSuite) TestService_PostUserDuplicateEmail() {
	path := suite.url + "/user/"
	user := &model.User{
		FirstName:   "Man",
		LastName:    "Lan",
		Email:       suite.duplicateEmail,
		CountryCode: "EN",
	}
	assert.Equal(suite.T(), http.StatusBadRequest, httpPost(path, user))
}

func (suite *ServiceTestSuite) TestService_PostUserInvalidData() {
	path := suite.url + "/user/"
	assert.Equal(suite.T(), http.StatusBadRequest, httpPost(path, []string{}))
}
