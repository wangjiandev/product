package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	configSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix), // 不设置前缀默认是/micro/config
		consul.StripPrefix(true),  // 是否移除前缀，直接获取配置
	)
	// 配置初始化
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	// 加载配置
	err = conf.Load(configSource)
	return conf, err
}
