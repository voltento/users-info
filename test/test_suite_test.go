package test_service

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/voltento/users-info/app/model"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(UsersInfoTestSuite))
}

type UsersInfoTestSuite struct {
	suite.Suite
	url      string
	testName string
	user1    *model.User
	user2    *model.User
}

func (s *UsersInfoTestSuite) cleanTestData() {
	_, users, err := s.GetUsers()
	if err != nil {
		assert.NoError(s.T(), err)
		return
	}

	for _, u := range users {
		s.DeleteUser(u.UserId)
	}
}

func (s *UsersInfoTestSuite) TearDownTest() {
	s.cleanTestData()
}

func (s *UsersInfoTestSuite) SetupTest() {
	s.url = "http://localhost:8181/"
	s.testName = "test_test_test"

	s.user1 = &model.User{
		FirstName:   s.testName,
		LastName:    "lastname1",
		Email:       "email1@test",
		CountryCode: "EN",
	}
	s.user2 = &model.User{
		FirstName:   s.testName,
		LastName:    "lastname2",
		Email:       "email2@test",
		CountryCode: "GB",
	}

	s.cleanTestData()
}

func (s *UsersInfoTestSuite) PostUser(user *model.User) int {
	data, err := json.Marshal(user)
	if err != nil {
		panic(err.Error())
	}

	req, err := http.NewRequest("POST", s.url+"user/", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func (s *UsersInfoTestSuite) DeleteUser(userId string) int {
	req, err := http.NewRequest("DELETE", s.url+"user/"+userId, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func (s *UsersInfoTestSuite) GetUsers(filters ...string) (int, []*model.User, error) {
	filter := ""
	for _, s := range filters {
		filter += "&" + s
	}

	req, err := http.NewRequest("GET", s.url+"users?first_name="+s.testName+filter, nil)
	if err != nil {
		return 0, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	var users []*model.User
	if err := json.Unmarshal(data, &users); err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, users, nil
}
