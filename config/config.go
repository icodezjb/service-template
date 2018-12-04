package config

import (
	"log"
	"sync"
)

var (
	configInstance *Config
)

type Config struct {
	Locker sync.RWMutex

	WebAddress string
}

func init() {
	configInstance = &Config{}
	if err := configInstance.load(); err != nil {
		panic("init config error: " + err.Error())
	}

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
