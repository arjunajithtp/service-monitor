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
	MonitoringIntervalInMin int      `json:"monitoringIntervalInMin"`
	DBHost                  string   `json:"dbHost"`
	DBPort                  string   `json:"dbPort"`
	DBName                  string   `json:"dbName"`
	DBUserName              string   `json:"dbUserName"`
	DBPassword              string   `json:"dbPassword"`
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
	return validateConfig()
}

func validateConfig() error {
	if Data.Port == "" {
		return fmt.Errorf("unable to find port in config file")
	}
	if Data.MonitoringIntervalInMin == 0 {
		return fmt.Errorf("unable to find monitoringIntervalInSec in config file")
	}
	if Data.DBHost == "" {
		return fmt.Errorf("unable to find dbHost in config file")
	}
	if Data.DBPort == "" {
		return fmt.Errorf("unable to find dbPort in config file")
	}
	if Data.DBName == "" {
		return fmt.Errorf("unable to find dbName in config file")
	}
	if Data.DBUserName == "" {
		return fmt.Errorf("unable to find dbUserName in config file")
	}
	if Data.DBPassword == "" {
		return fmt.Errorf("unable to find dbPassword in config file")
	}
	if len(Data.Services) == 0 {
		return fmt.Errorf("unable to find services in config file")
	}
	return nil
}
