package config

type DatarefConfig struct {
	Name       string `yaml:"name"`
	DatarefStr string `yaml:"value"`
}
type ServerConfig struct{}

type Config struct {
	DatarefConfig []DatarefConfig
	ServerConfig  ServerConfig
}

const ServerUrl = "https://api.xairline.org"
