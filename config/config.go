package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Configs struct {
	PublicAddr      string `yaml:"public_address"`
	Port            int    `yaml:"port"`
	MySQLDSN        string `yaml:"mysql_dsn"`
	LocalVideoPath  string `yaml:"local_video_path"`
	RemoteVideoPath string `yaml:"remote_video_path"`
	Url             string
}

var Config Configs

func init() {
	f, err := os.OpenFile("./config.yaml", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	// read yaml into Config
	err = yaml.NewDecoder(f).Decode(&Config)
	if err != nil {
		panic(err)
	}

	Config.Url = fmt.Sprintf("http://%s:%d/", Config.PublicAddr, Config.Port)
}
