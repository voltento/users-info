package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigFromBytesDataBase(t *testing.T) {
	cfgData := `
{
  "database": {
    "user": "test1",
    "password": "test2",
    "database": "test3"
  }
}	
`

	config, err := NewConfigFromBytes([]byte(cfgData))

	assert.NoError(t, err)

	db := config.Database
	assert.Equal(t, db.User, "test1", "db.user")
	assert.Equal(t, db.Password, "test2", "db.password")
	assert.Equal(t, db.Database, "test3", "db.database")
}

func TestNewConfigFromBytesLogger(t *testing.T) {
	cfgData := `
{
  "logger": {
    "level" :"error"
  }
}	
`

	config, err := NewConfigFromBytes([]byte(cfgData))

	assert.NoError(t, err)

	logger := config.Logger
	assert.Equal(t, logger.Level, "error", "logger.level")
}

func TestNewConfigFromBytesService(t *testing.T) {
	cfgData := `
{
  "service": {
    "address" :"mail:221",
	"log_gin_gonic": true
  }
}	
`

	config, err := NewConfigFromBytes([]byte(cfgData))
	assert.NoError(t, err)

	service := config.Service

	assert.Equal(t, service.Address, "mail:221", "service.address")
	assert.Equal(t, service.LogGinGonic, true, "service.logGinGonic")
}

func TestNewConfigFromBytesInvaludJson(t *testing.T) {
	cfgData := `
{
  "service": {
}	
`
	_, err := NewConfigFromBytes([]byte(cfgData))
	assert.Error(t, err)

}
