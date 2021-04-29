/**
 *@Description
 *@ClassName plugin
 *@Date 2020/11/24 1:38 下午
 *@Author ckhero
 */

package plugin

import (
	"fmt"
	"github.com/ckhero/go-common/config"
	"google.golang.org/grpc"
)

type Plugin interface {
	InitPlugin() error
	Release()
}

type PluginCtx struct {
	plugins []Plugin
	Errc chan error
}

func NewPluginCtx(errc chan error, plugins ...Plugin) *PluginCtx {

	defaultPlugins := []Plugin{NewPluginLogger()}

	plugins = append(defaultPlugins, plugins...)

	return &PluginCtx{
		plugins: plugins,
		Errc:    errc,
	}
}

func (ms *PluginCtx) InitPlugin() {
	for _, p := range ms.plugins {
		err := p.InitPlugin()
		if err != nil {
			panic(fmt.Errorf("Fatal init server: %s \n", err))
		}
	}
}

func (ms *PluginCtx) Release()  {
	for _, p := range ms.plugins {
		go p.Release()
	}
}

func (ms *PluginCtx) HealthCheck(s *grpc.Server)  {
}

func GetPlugins(srvName string) []Plugin {
	defaultPlugins := []Plugin{
		NewPluginMysql(),
		NewPluginRedis(),
	}
	for _, plugin := range config.GetPluginsCfg() {
		switch plugin {
		}
	}
	return defaultPlugins
}

