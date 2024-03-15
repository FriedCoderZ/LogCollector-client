package config

import (
	"log"

	"github.com/spf13/viper"
)

// 读取配置文件config
type Config struct {
	Crypto    CryptoConfig
	Server    ServerConfig
	Collector CollectorConfig
}

var (
	config Config
)

func init() {
	// 把配置文件读取到结构体上
	viper.SetConfigName("config")
	// viper.AddConfigPath("../../")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	viper.Unmarshal(&config) //将配置文件绑定到config上
}

func GetConfig() Config {
	return config
}
