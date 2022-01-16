package config

import (
	"io/ioutil"
	"xairline/goxairline/internal/xplane/shared"

	"github.com/xairline/goplane/extra/logging"
	"gopkg.in/yaml.v2"
)

func NewConfig(configFile string) *Config {
	var res Config
	var err error

	res.DatarefConfig, err = getDatarefConfig(configFile, nil)
	if err != nil {
		logging.Errorf("Failed to load dataref config: %s", configFile)
		return nil
	}

	res.ServerConfig, err = getServerConfig(SERVER_URL)
	if err != nil {
		logging.Errorf("Failed to load dataref config: %s", configFile)
		return nil
	}

	return &res
}

func getServerConfig(SERVER_URL string) (ServerConfig, error) {
	logging.Error("unimplemented")
	return ServerConfig{}, nil
}

func getDatarefConfig(configFile string, logger *shared.Logger) ([]DatarefConfig, error) {
	var res []DatarefConfig
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		logger.Errorf("Failed to get yaml file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &res)
	if err != nil {
		logger.Errorf("Unmarshal: %v", err)
	}
	logger.Infof("Get config object from file: %v", res)
	return res, nil
}
