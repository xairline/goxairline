package config

import "github.com/xairline/goplane/extra/logging"

func NewConfig(configFile string) *Config {
	var res Config
	var err error

	res.DatarefConfig, err = getDatarefConfig(configFile)
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

func getDatarefConfig(configFile string) (DatarefConfig, error) {
	logging.Error("unimplemented")
	return DatarefConfig{}, nil
}
