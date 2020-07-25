package config

import (
	"encoding/json"
	"io/ioutil"
)

func NewConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewConfigFromBytes(data)
}

func NewConfigFromBytes(data []byte) (*Config, error) {
	config := NewDefaultConfig()
	err := json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, err
}
