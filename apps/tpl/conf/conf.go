package conf

import (
	"fmt"
	"log"
	"sync"

	"gitee.com/ha666/golibs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	_key = "?GQ$0K0GgLdO=f+~L68PLm$uhKr4'=tV"
	_iv  = "VFs7@sK61cj^f?HZ"
)

func init() {
	Parse("/etc/gourd/config")
}

// config 全部配置
type config struct {
	// DNS 认证配置
	DNS struct {
		ID      string   `mapstructure:"id"`
		Token   string   `mapstructure:"token"`
		Domain  string   `mapstructure:"domain"`
		KeyWord string   `mapstructure:"keyword"`
		Names   []string `mapstructure:"names"`
	} `mapstructure:"dnspod"`

	// Gitlab 配置
	Gitlab struct {
		BaseURL     string `mapstructure:"base_url"`
		AccessToken string `mapstructure:"access_token"`
	} `mapstructure:"gitlab"`

	// MySQL mysql配置
	DB struct {
		DSN     string `mapstructure:"dsn"`
		MaxConn int    `mapstructure:"maxconn"`
	} `mapstructure:"db"`
}

// Config 配置单例
var (
	Config config
	// hot update need lock
	mu sync.RWMutex
)

func (c *config) Set(cp *config) {
	mu.Lock()
	defer mu.Unlock()
	*c = *cp
}

func (c *config) GetDNSID() string {
	mu.RLock()
	defer mu.RUnlock()
	return c.DNS.ID
}

func (c *config) GetGitlabBaseURL() string {
	mu.RLock()
	defer mu.RUnlock()
	return c.Gitlab.BaseURL
}

func (c *config) GetGitlabToken() string {
	mu.RLock()
	defer mu.RUnlock()
	return c.Gitlab.AccessToken
}

func (c *config) GetDBDSN() string {
	mu.RLock()
	defer mu.RUnlock()
	return c.DB.DSN
}

func (c *config) GetDBMaxConn() int {
	mu.RLock()
	defer mu.RUnlock()
	return c.DB.MaxConn
}

// Parse 配置解析
func Parse(path string) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := parseConfig(); err != nil {
		log.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		if event.Op == fsnotify.Write {
			if err := parseConfig(); err != nil {
				log.Printf("hot update failed, %s, use last config", err.Error())
			}
		}
	})
}

func parseConfig() error {
	var c config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("viper unmarshal failed")
	}

	decode := func(str *string) error {
		if *str == "" {
			return fmt.Errorf("decode field is empty string")
		}
		var tmpBytes []byte
		log.Printf("%s,%s,%s", *str, _key, _iv)
		tmpBytes, err = golibs.AesDecrypt(golibs.HexStringToBytes(*str), []byte(_key), []byte(_iv))
		if err != nil {
			log.Println(err)
			return fmt.Errorf("AES Decrypt failed")
		}
		*str = string(tmpBytes)
		return nil
	}

	if err := decode(&c.DB.DSN); err != nil {
		return err
	}

	Config.Set(&c)
	return nil
}
