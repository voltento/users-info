package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/voltento/users-info/app/connectors/storage"
	"github.com/voltento/users-info/app/fault"
	"github.com/voltento/users-info/app/logger"
	"github.com/voltento/users-info/app/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite
	url                 string
	service             *Service
	modelUserToUserData map[model.User][]model.User
	userIdToUserData    map[string]*model.User
	testUser1           model.User
	testUser2           model.User
	duplicateEmail      string
}

func (suite *ServiceTestSuite) users(modelUser *model.User) ([]model.User, error) {
	if len(suite.modelUserToUserData) == 0 {
		return nil, fault.NewNotFound("no users")
	}
	if v, isOk := suite.modelUserToUserData[*modelUser]; isOk {
		return v, nil
	}
	return nil, fault.NewNotFound("no users")
}

func (suite *ServiceTestSuite) user(userId string) (*model.User, error) {
	if len(userId) == 0 {
		return nil, fault.NewBadRequest("empty user id")
	}

	if _, err := strconv.Atoi(userId); err != nil {
		return nil, fault.NewBadRequest("empty user id")
	}

	if u, isOk := suite.userIdToUserData[userId]; isOk {
		return u, nil
	}
	return nil, fault.NewNotFound(fmt.Sprintf("no user with id '%v'", userId))
}

func (suite *ServiceTestSuite) deleteUser(userId string) error {
	if len(userId) == 0 {
		return fault.NewBadRequest("received empty user id")
	}

	if _, err := strconv.Atoi(userId); err != nil {
		return fault.NewBadRequest("received invalud user id")
	}

	if _, isOk := suite.userIdToUserData[userId]; !isOk {
		return fault.NewNotFound("user not found")
	} else {
		delete(suite.userIdToUserData, userId)
		return nil
	}
}

func (suite *ServiceTestSuite) addUser(u *model.User) error {
	isBadRequest := false
	isBadRequest = isBadRequest || len(u.FirstName) == 0
	isBadRequest = isBadRequest || len(u.LastName) == 0
	isBadRequest = isBadRequest || len(u.Email) == 0
	isBadRequest = isBadRequest || len(u.CountryCode) == 0

	if isBadRequest {
		return fault.NewBadRequest("bad request")
	}

	if u.Email == suite.duplicateEmail {
		return fault.NewBadRequest("email already exists")
	}

	return nil
}

func (suite *ServiceTestSuite) updateUser(u *model.User) error {
	isBadRequest := false
	isBadRequest = isBadRequest || len(u.FirstName) == 0
	isBadRequest = isBadRequest || len(u.LastName) == 0
	isBadRequest = isBadRequest || len(u.Email) == 0
	isBadRequest = isBadRequest || len(u.CountryCode) == 0

	if isBadRequest {
		return fault.NewBadRequest("bad request")
	}

	return nil
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.service.Stop()
}

func (suite *ServiceTestSuite) SetupTest() {
	config := &Config{
		Address:     "localhost:8181",
		LogGinGonic: false,
	}
	suite.url = "http://" + config.Address

	s := storage.NewStorageMock(suite.users, suite.user, suite.deleteUser, suite.updateUser, suite.addUser)

	err, service := NewService(config, logger.NewMock().Sugar(), s)
	if err != nil {
		panic(err.Error())
	}
	suite.service = service

	suite.modelUserToUserData = make(map[model.User][]model.User)
	suite.userIdToUserData = make(map[string]*model.User)

	suite.testUser1 = model.User{
		UserId:      "1",
		FirstName:   "testUser1_firstname",
		LastName:    "testUser1_LastName",
		Email:       "testUser1_email",
		CountryCode: "testUser1_countrycode",
	}
	suite.testUser2 = model.User{
		UserId:      "2",
		FirstName:   "testUser2_firstname",
		LastName:    "testUser2_LastName",
		Email:       "testUser2_email",
		CountryCode: "testUser2_countrycode",
	}

	suite.duplicateEmail = "exists@mail"

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		go suite.service.Run()
		wait := true
		time.AfterFunc(time.Second*5, func() {
			wait = false
		})
		for wait {
			if resp, er := http.Get(suite.url + "/healthcheck"); er == nil {
				resp.Body.Close()
				return
			}
		}
	}()
	wg.Wait()
}

func (suite *ServiceTestSuite) checkUsersFromHttpResponse(expect []model.User, resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	var user []model.User
	err = json.Unmarshal(data, &user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(expect), len(user))
	assert.Equal(suite.T(), usersToSet(expect), usersToSet(user))
}

func (suite *ServiceTestSuite) checkUserFromHttpResponse(expect model.User, resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	var user model.User
	err = json.Unmarshal(data, &user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, user)
}

func usersToSet(users []model.User) map[model.User]struct{} {
	result := make(map[model.User]struct{}, len(users))
	for _, u := range users {
		if _, hasVal := result[u]; hasVal {
			panic(fmt.Sprintf("usersToSet: get duplicate values: %v", u))
		}
		result[u] = struct{}{}
	}
	return result
}

func httpDelete(host string) int {
	req, err := http.NewRequest("DELETE", host, nil)
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

func httpPut(host string) int {
	req, err := http.NewRequest("PUT", host, nil)
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

func httpPost(host string, obj interface{}) int {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err.Error())
	}

	req, err := http.NewRequest("POST", host, bytes.NewReader(data))
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
