/**
 *@Description
 *@ClassName wx
 *@Date 2021/5/13 上午9:26
 *@Author ckhero
 */

package wx

import (
	"github.com/ckhero/go-common/config"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniCfg "github.com/silenceper/wechat/v2/miniprogram/config"
	wechatCache "github.com/silenceper/wechat/v2/cache"
	"sync"
)

var onceMiniPay sync.Once
var onceMini sync.Once
var miniPayClient *wechat.Client
var miniClient *miniprogram.MiniProgram

func newMiniPayClient(cfg *config.WeixinPay) *wechat.Client{
	onceMiniPay.Do(func() {
		mchId := cfg.MchId
		appId := cfg.AppIdMini
		apiKey := cfg.ApiKey
		isProd := cfg.IsProd
		miniPayClient = wechat.NewClient(appId, mchId, apiKey, isProd)
	})
	return miniPayClient
}

func NewMiniPayClient(cfg *config.WeixinPay) *wechat.Client{
	return newMiniPayClient(config.GetWeixinPayCfg())
}

// 获取小程序的单例
func NewMiniClient() (*miniprogram.MiniProgram) {

	onceMini.Do(func() {
		cfg := config.GetWeixinCfg()
		miniClient = miniprogram.NewMiniProgram(&miniCfg.Config{
			AppID: cfg.AppId,
			AppSecret: cfg.SecretKey,
			Cache: &wechatCache.Memory{},
		})
	})

	return miniClient
}
