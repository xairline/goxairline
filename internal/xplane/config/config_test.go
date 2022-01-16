package config

import (
	"os"
	"reflect"
	"testing"
	"xairline/goxairline/internal/xplane/shared"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.configFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getServerConfig(t *testing.T) {
	type args struct {
		SERVER_URL string
	}
	tests := []struct {
		name    string
		args    args
		want    ServerConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getServerConfig(tt.args.SERVER_URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("getServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDatarefConfig(t *testing.T) {
	path, _ := os.Getwd()
	t.Logf("Current test filename: %s", path)
	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		want    []DatarefConfig
		wantErr bool
	}{
		{
			name: "should load config file",
			args: args{configFile: path + "/config.yaml"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := shared.GetLoggerForTest(t)
			got, err := getDatarefConfig(tt.args.configFile, &logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDatarefConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("getDatarefConfig failed, error: %+v", err)
			}
		})
	}
}
