package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Data will have the configuration data from config file
var Data configModel

// configModel is for holding all the run time configurations for the application
type configModel struct {
	Port                    string   `json:"port"`
	MonitoringIntervalInSec int      `json:"monitoringIntervalInSec"`
	Services                []string `json:"services"`
}

// SetConfiguration will extract the config data from file
func SetConfiguration() error {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		return fmt.Errorf("error: no value set to environment variable CONFIG_FILE_PATH")
	}
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error: failed to read the config file: %v", err)
	}
	err = json.Unmarshal(configData, &Data)
	if err != nil {
		return fmt.Errorf("error: failed to unmarshal config data: %v", err)
	}
	return nil
}
