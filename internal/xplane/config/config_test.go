package config

import (
	"os"
	"reflect"
	"testing"
	"xairline/goxairline/internal/xplane/shared"
)

func TestNewConfig(t *testing.T) {
	path, _ := os.Getwd()
	type args struct {
		configFile string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "should not load non-exist config file",
			args: args{configFile: path + "/config.yaml2"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := shared.GetLoggerForTest(t)
			if got := NewConfig(tt.args.configFile, &logger); !reflect.DeepEqual(got, tt.want) {
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
		{
			name:    "should load remote config file",
			args:    args{SERVER_URL: "TBD"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := shared.GetLoggerForTest(t)
			got, err := getServerConfig(tt.args.SERVER_URL, &logger)
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
			name:    "should load config file",
			args:    args{configFile: path + "/config.yaml"},
			wantErr: false,
		},
		{
			name:    "should load config file that doesn't exist",
			args:    args{configFile: path + "/config.yaml2"},
			want:    nil,
			wantErr: true,
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
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
