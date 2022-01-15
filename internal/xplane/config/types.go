package config

type DatarefConfig struct {
	Name        string
	DatarefStr  string
	DatarefType string
}
type ServerConfig struct{}

type Config struct {
	DatarefConfig DatarefConfig
	ServerConfig  ServerConfig
}

const SERVER_URL = "https://api.xairline.org"
