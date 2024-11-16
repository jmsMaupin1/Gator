package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL string            `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", homeDir, configFileName), nil
}

func (c *Config) write() error {
	configFilePath, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("Error getting config path: %v", err)
	}

	data, err := json.Marshal(*c)
	if err != nil {
		return fmt.Errorf("Json marshalling error while writing: %v", err)
	}

	if err := os.WriteFile(configFilePath, data, 0666); err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

func Read() (Config, error){
	filePath, err := getConfigPath()
	if err != nil {
		return Config{}, fmt.Errorf("Error retrieving config path: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("file reading error: %v", err)
	}

	var config Config	
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("Unmarshalling error: %v", err)
	}

	return config, nil
}

func (c *Config) SetUser(user_name string) error {
	c.CurrentUserName = user_name
	if err := c.write(); err != nil {
		return err
	}

	return nil
}
