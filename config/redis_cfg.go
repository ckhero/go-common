/**
 *@Description
 *@ClassName redis
 *@Date 2021/4/29 下午2:59
 *@Author ckhero
 */

package config


type Redis struct {
	Host        string `json:"host" yaml:"host"`
	Port        uint16 `json:"port" yaml:"port"`
	Password    string `json:"password" yaml:"password"`
	Database    uint8  `json:"database" yaml:"database"`
	MaxIdle     int    `json:"maxIdle" yaml:"maxIdle"`         // 空闲连接数大小
	MaxActive   int    `json:"maxActive" yaml:"maxActive"`     // 最大连接数
	IdleTimeout int    `json:"idleTimeout" yaml:"idleTimeout"` //空闲连接超时时间
	Prefix      string `json:"prefix" yaml:"prefix"`           //前缀使用系统名字缩写
}


func GetRedisCfg() map[string]Redis {
	return appConfig.Redis
}