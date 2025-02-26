package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Token string `json:"token"`
	Proxy string `json:"proxy"`
}

func ReadOrCreateConfig(filePath string) (config *Config, err error) {
	// check file
	_, err = os.Stat(filePath)
	emptyErr := fmt.Errorf("enter your token in %v", filePath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		// if the file does not exist, a default config is created
		defaultConfig := &Config{}

		configData, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			err = fmt.Errorf("failed to marshal default config: %w", err)
			return nil, err
		}
		if err = os.WriteFile(filePath, configData, 0644); err != nil {
			err = fmt.Errorf("failed to write config file: %w", err)
			return nil, err
		}
		return nil, emptyErr
	}

	configData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// parse json
	err = json.Unmarshal(configData, &config)
	if err != nil {
		err = fmt.Errorf("parsing config file failed: %w", err)
	}

	//check
	if len(config.Token) == 0 {
		err = emptyErr
	}

	return
}
