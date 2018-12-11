package conf

import (
	"log"

	"gitee.com/ha666/golibs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	Parse("/etc/gourd/config", &Config)
}

const (
	_key = "?GQ$0K0GgLdO=f+~L68PLm$uhKr4'=tV"
	_iv  = "VFs7@sK61cj^f?HZ"
)

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

	if err := parseConfig(model); err != nil {
		log.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Op == fsnotify.Write {
			if err := parseConfig(model); err != nil {
				log.Fatal(err)
			}
		}
	})
}

func parseConfig(model interface{}) error {
	err := viper.Unmarshal(model)
	if err != nil {
		log.Fatal(err)
	}

	decode := func(str *string) {
		if *str == "" {
			return
		}
		var tmpBytes []byte
		log.Printf("%s,%s,%s", *str, _key, _iv)
		tmpBytes, err = golibs.AesDecrypt(golibs.HexStringToBytes(*str), []byte(_key), []byte(_iv))
		if err != nil {
			return
		}
		*str = string(tmpBytes)
	}

	decode(&Config.DB.DSN)
	return nil
}
