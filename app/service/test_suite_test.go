package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/voltento/users-info/app/connectors/storage"
	"github.com/voltento/users-info/app/logger"
	"github.com/voltento/users-info/app/model"
	"io/ioutil"
	"net/http"
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
}

func (suite *ServiceTestSuite) usersFunc(modelUser *model.User) ([]model.User, error) {
	if len(suite.modelUserToUserData) == 0 {
		return nil, errors.New("map is empty")
	}
	if v, isOk := suite.modelUserToUserData[*modelUser]; isOk {
		return v, nil
	}
	return nil, errors.New("no users")
}

func (suite *ServiceTestSuite) userFunc(userId string) (*model.User, error) {
	if len(userId) == 0 {
		return nil, errors.New("received empty user id")
	}

	if u, isOk := suite.userIdToUserData[userId]; isOk {
		return u, nil
	}
	return nil, errors.New(fmt.Sprintf("no user with id '%v'", userId))
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

	s := storage.NewStorageMock(suite.usersFunc, suite.userFunc)

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
		SecondName:  "testUser1_secondname",
		Email:       "testUser1_email",
		CountryCode: "testUser1_countrycode",
	}
	suite.testUser2 = model.User{
		UserId:      "2",
		FirstName:   "testUser2_firstname",
		SecondName:  "testUser2_secondname",
		Email:       "testUser2_email",
		CountryCode: "testUser2_countrycode",
	}

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

func (suite *ServiceTestSuite) checkUsersFromHttpResponse(expect model.User, resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	var user []model.User
	err = json.Unmarshal(data, &user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(user))
	if len(user) > 0 {
		assert.Equal(suite.T(), expect, user[0])
	}
}

func (suite *ServiceTestSuite) checkUserFromHttpResponse(expect model.User, resp *http.Response) {
	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	var user model.User
	err = json.Unmarshal(data, &user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expect, user)
}
