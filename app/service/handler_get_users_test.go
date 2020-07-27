package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/voltento/users-info/app/model"
	"net/http"
)

func (suite *ServiceTestSuite) TestService_GetUsersGetByUserIdOk() {
	suite.modelUserToUserData[model.User{UserId: "1"}] = []model.User{
		suite.testUser1,
	}

	query := suite.url + "/users?user_id=1"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUsersFromHttpResponse([]model.User{suite.testUser1}, resp)
}

func (suite *ServiceTestSuite) TestService_GetUsersGetByUserIdNotFound() {
	query := suite.url + "/users?user_id=1"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), resp.StatusCode, http.StatusBadRequest)
}

func (suite *ServiceTestSuite) TestService_GetUsersGetByFirstNameOk() {
	suite.modelUserToUserData[model.User{FirstName: "aa"}] = []model.User{
		suite.testUser1,
	}

	query := suite.url + "/users?first_name=aa"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUsersFromHttpResponse([]model.User{suite.testUser1}, resp)
}

func (suite *ServiceTestSuite) TestService_GetUsersGetByFirstNameTwoUsersOk() {
	var testUser1Copy model.User
	testUser1Copy = suite.testUser1
	testUser1Copy.UserId = "2"
	respUsers := []model.User{
		suite.testUser1,
		testUser1Copy,
	}
	suite.modelUserToUserData[model.User{FirstName: "aa"}] = respUsers
	query := suite.url + "/users?first_name=aa"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUsersFromHttpResponse(respUsers, resp)
}

func (suite *ServiceTestSuite) TestService_GetUsersGetBySecondNameOk() {
	suite.modelUserToUserData[model.User{SecondName: "second_name"}] = []model.User{
		suite.testUser1,
	}

	query := suite.url + "/users?second_name=second_name"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUsersFromHttpResponse([]model.User{suite.testUser1}, resp)
}

func (suite *ServiceTestSuite) TestService_GetUsersGetByEmailOk() {
	suite.modelUserToUserData[model.User{Email: "foo@mail.com"}] = []model.User{
		suite.testUser1,
	}

	query := suite.url + "/users?email=foo@mail.com"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUsersFromHttpResponse([]model.User{suite.testUser1}, resp)
}

func (suite *ServiceTestSuite) TestService_GetUsersGetByCountryOk() {
	suite.modelUserToUserData[model.User{CountryCode: "EN"}] = []model.User{
		suite.testUser1,
	}

	query := suite.url + "/users?country_code=EN"
	resp, err := http.Get(query)

	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	assert.Equal(suite.T(), resp.StatusCode, http.StatusOK)
	suite.checkUsersFromHttpResponse([]model.User{suite.testUser1}, resp)
}
