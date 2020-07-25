package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func NewConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath.Clean(path))
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
