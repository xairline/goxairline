package config

import (
	"io/ioutil"
	"xairline/goxairline/internal/xplane/shared"

	"gopkg.in/yaml.v2"
)

func NewConfig(configFile string, logger *shared.Logger) *Config {
	var res Config
	var err error

	res.DatarefConfig, err = getDatarefConfig(configFile, logger)
	if err != nil {
		logger.Errorf("Failed to load dataref config: %s", configFile)
		return nil
	}

	res.ServerConfig, err = getServerConfig(ServerUrl, logger)
	if err != nil {
		logger.Errorf("Failed to load dataref config: %s", configFile)
		return nil
	}

	return &res
}

func getServerConfig(SERVER_URL string, logger *shared.Logger) (ServerConfig, error) {
	logger.Errorf("unimplemented")
	return ServerConfig{}, nil
}

func getDatarefConfig(configFile string, logger *shared.Logger) ([]DatarefConfig, error) {
	var res []DatarefConfig
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		logger.Errorf("Failed to get yaml file: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &res)
	if err != nil {
		logger.Errorf("Unmarshal: %v", err)
		return nil, err
	}
	logger.Infof("Get config object from file: %v", res)
	return res, nil
}
