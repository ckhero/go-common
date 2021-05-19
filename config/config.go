/**
 *@Description
 *@ClassName config
 *@Date 2020/11/23 2:16 下午
 *@Author ckhero
 */

package config

import (
	"fmt"
	sysConst "github.com/ckhero/go-common/constant/sys"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"sync"
)

type Config struct {
	Project        string              `yaml:"project"`
	Application    string              `yaml:"application"`
	Env            string              `yaml:"env"`
	Debug          ServerInfo          `yaml:"debug"`
	HTTP           ServerInfo          `yaml:"http"`
	GRPC           ServerInfo          `yaml:"grpc"`
	Database       map[string]Database `yaml:"database"`
	Redis          map[string]Redis    `yaml:"redis"`
	Domain         map[string]string   `yaml:"domain"`
	Logger         Logger              `yaml:"logger"`
	Plugins        []string            `yaml:"plugins"`
	WeixinPay      *WeixinPay          `yaml:"weixinPay"`
	Weixin         *Weixin             `yaml:"weixin"`
	Auth           *Auth               `yaml:"auth"`
	RedPacketLimit *RedPacketLimit     `yaml:"redPacketLimit"`
}

type ServerInfo struct {
	Host string
	Port int
}

type Logger struct {
	Level string
}

var (
	appConfig  *Config
	configOnce sync.Once
)

func InitConfig(path string) *Config {
	configOnce.Do(func() {
		appConfig = &Config{}
		commonPath := strings.ReplaceAll(path, filepath.Base(path), "common.yaml")
		loadConfig(appConfig, path)
		loadConfig(appConfig, commonPath)
	})
	return appConfig
}

func loadConfig(cfg *Config, path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(cfg) // 将配置信息绑定到结构体上
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error unmarshal config file %s \n", err))
	}
}
func GetGlobalCfg() *Config {
	return appConfig
}

func GetEnv() string {
	if appConfig.Env == "" {
		return sysConst.SysEnvTest
	}
	return appConfig.Env
}

func GetPluginsCfg() []string {
	return appConfig.Plugins
}
