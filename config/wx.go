/**
 *@Description
 *@ClassName wx
 *@Date 2021/5/13 上午9:27
 *@Author ckhero
 */

package config

type WeixinPay struct {
	MchId     string `yaml:"mchId"`
	ApiKey    string `yaml:"apiKey"`
	IsProd    bool   `yaml:"isProd"`
	NotifyUrl string `yaml:"notifyUrl"`
	AppIdMini string `yaml:"appIdMini"`
	AppIdApp  string `yaml:"appIdApp"`

	// 3个证书路径
	CertPem string `yaml:"certPem"`
	KeyPem  string `yaml:"keyPem"`
	CertP12 string `yaml:"certP12"`
}

type Weixin struct {
	AppId string `yaml:"appId"`
	SecretKey string `yaml:"secretKey"`
}

func GetWeixinPayCfg() *WeixinPay {
	return appConfig.WeixinPay
}

func GetWeixinCfg() *Weixin {
	return appConfig.Weixin
}