package database

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/voltento/users-info/app/fault"
	"github.com/voltento/users-info/app/model"
	"testing"
)

func TestConfigToPgOptionsOk(t *testing.T) {
	cfg := &Config{
		User:     "test_user",
		Password: "test_password",
		Database: "test_database",
		Addr:     "local_host:8080",
	}

	pgOptions, err := configToPgOptions(cfg)

	assert.NoError(t, err)
	assert.Equal(t, cfg.User, pgOptions.User)
	assert.Equal(t, cfg.Password, pgOptions.Password)
	assert.Equal(t, cfg.Database, pgOptions.Database)
	assert.Equal(t, cfg.Addr, pgOptions.Addr)
}

func TestConfigToPgOptionsNoUser(t *testing.T) {
	cfg := &Config{
		Password: "test_password",
		Database: "test_database",
	}

	_, err := configToPgOptions(cfg)
	assert.Error(t, err)
}

func TestConfigToPgOptionsNoPassword(t *testing.T) {
	cfg := &Config{
		User:     "test_user",
		Database: "test_database",
	}

	_, err := configToPgOptions(cfg)
	assert.Error(t, err)
}

func TestConfigToPgOptionsNoDatabase(t *testing.T) {
	cfg := &Config{
		User:     "test_user",
		Password: "test_password",
	}

	_, err := configToPgOptions(cfg)
	assert.Error(t, err)
}

func TestDtoUserToModelUser(t *testing.T) {
	dtoUser := &User{
		UserId:      1,
		FirstName:   "test_firstname",
		LastName:    "test_lastname",
		Email:       "test_email",
		CountryCode: "test_country_code",
	}

	modelUser := dtoUserToModelUser(dtoUser)
	assert.Equal(t, modelUser.UserId, "1")
	assert.Equal(t, modelUser.FirstName, dtoUser.FirstName)
	assert.Equal(t, modelUser.LastName, dtoUser.LastName)
	assert.Equal(t, modelUser.Email, dtoUser.Email)
	assert.Equal(t, modelUser.CountryCode, dtoUser.CountryCode)
}

func TestModelUserToDtoUserOk(t *testing.T) {
	modelUser := &model.User{
		UserId:      "2",
		FirstName:   "test_firstname",
		LastName:    "test_LastName",
		Email:       "test_email@fo",
		CountryCode: "test_country_code",
	}

	dtoUser, err := modelUserToDtoUser(modelUser)

	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Equal(t, dtoUser.UserId, 2)
	assert.Equal(t, dtoUser.FirstName, modelUser.FirstName)
	assert.Equal(t, dtoUser.LastName, modelUser.LastName)
	assert.Equal(t, dtoUser.Email, modelUser.Email)
	assert.Equal(t, dtoUser.CountryCode, modelUser.CountryCode)
}

func TestModelUserToDtoUserWrongEmail(t *testing.T) {
	modelUser := &model.User{Email: "test_email"}
	_, err := modelUserToDtoUser(modelUser)
	assert.Error(t, err)
}

func TestModelUserToDtoNoUserId(t *testing.T) {
	modelUser := &model.User{}

	dtoUser, err := modelUserToDtoUser(modelUser)
	assert.NoError(t, err)
	assert.Equal(t, dtoUser.UserId, 0)
}

func TestModelUserToDtoUserWrongUserId(t *testing.T) {
	modelUser := &model.User{
		UserId: "a",
	}

	_, err := modelUserToDtoUser(modelUser)
	assert.Error(t, err)
}

func TestConfigToPgOptionsEmpty(t *testing.T) {
	_, err := configToPgOptions(nil)
	assert.Error(t, err)
}

func Test_sqlErrorToErrorColumnViolationBadRequest(t *testing.T) {
	err := errors.New("violates not-null constraint in column")
	_, isOk := sqlErrorToError(err).(*fault.BadRequest)
	assert.True(t, isOk)
}

func Test_sqlErrorToErrorColumnViolation(t *testing.T) {
	err := errors.New("some error")
	_, isOk := sqlErrorToError(err).(*fault.BadRequest)
	assert.False(t, isOk)
}
