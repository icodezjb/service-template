package config

import (
	"log"
	"sync"

	"github.com/buchenglei/service-template/common/definition"
)

var (
	configInstance *Config

	ServiceEnv definition.Environment
)

type Config struct {
	Locker sync.RWMutex

	Env        string
	WebAddress string
}

func init() {
	configInstance = &Config{}
	if err := configInstance.load(); err != nil {
		panic("init config error: " + err.Error())
	}

	// 程序启动后初始化一次环境变量
	ServiceEnv = definition.Environment(configInstance.Env)

	configInstance.watch()
}

func Get() *Config {
	return configInstance
}

func (c *Config) load() error {
	c.Locker.Lock()
	// 加载所有配置项

	c.Locker.Unlock()
	return nil
}

func (c *Config) watch() {
	// 监听配置更新
	// update config
	go func() {
		err := c.load()
		if err != nil {
			log.Println("update config error: %v", err)
		}
	}()
}
