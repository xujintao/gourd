package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	Parse("/etc/gourd/config", &Config)
}

// config 全部配置
type config struct {
	// Auth 认证配置
	Auth struct {
		JWT struct {
			Secret string `mapstructure:"secret"`
			AppID  int    `mapstructure:"appid"`
		} `mapstructure:"jwt"`
	} `mapstructure:"auth"`

	// MySQL mysql配置
	DB struct {
		DSN     string `mapstructure:"dsn"`
		MaxConn int    `mapstructure:"maxconn"`
	} `mapstructure:"db"`
}

// Config 配置单例
var Config config

// Parse 配置解析
func Parse(path string, model interface{}) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	err := viper.Unmarshal(model)
	if err != nil {
		log.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Op == fsnotify.Write {
			err := viper.Unmarshal(model)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}
