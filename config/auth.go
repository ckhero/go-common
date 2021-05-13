/**
 *@Description
 *@ClassName auth
 *@Date 2021/5/13 下午2:18
 *@Author ckhero
 */

package config


type Auth struct {
	SecretKey string `yaml:"secretKey"`
}

func GetAuthCfg() *Auth {
	return appConfig.Auth
}