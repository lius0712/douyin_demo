package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Configs struct {
	PublicAddr      string `yaml:"public_address"`
	DbPort          int    `yaml:"db_port"`
	MySQLDSN        string `yaml:"mysql_dsn"`
	LocalVideoPath  string `yaml:"local_video_path"`
	RemoteVideoPath string `yaml:"remote_video_path"`
	RedisAddr       string `yaml:"redis_address"`
	RdbPort         int    `yaml:"rdb_port"`
	RdbPwd          string `yaml:"rdb_pwd"`
	RdbNum          int    `yaml:"rdb_num"`
	RdbPoolSize     int    `yaml:"rdb_pool_size"`
	DbUrl           string
	RdbUrl          string
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

	Config.DbUrl = fmt.Sprintf("http://%s:%d/", Config.PublicAddr, Config.DbPort)
	Config.RdbUrl = fmt.Sprintf("%s:%d", Config.RedisAddr, Config.RdbPort)
}
