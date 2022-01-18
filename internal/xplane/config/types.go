package config

type DatarefConfig struct {
	Name         string `yaml:"name"`
	DatarefStr   string `yaml:"value"`
	Precision    int    `yaml:"precision,omitempty"`
	IsBytesArray bool   `yaml:"isBytesArray,omitempty"`
}
type ServerConfig struct{}

type Config struct {
	DatarefConfig []DatarefConfig
	ServerConfig  ServerConfig
}

const SERVER_URL = "https://api.xairline.org"
