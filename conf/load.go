package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 配置映射成config对象

func LoadConfigFormToml(filePath string) error {
	config = NewDefaultConfig()
	// 读取配置文件
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config file error,path: %s,%s", err, filePath)
	}
	return loadGlobal()
}
func LoadConfigFormEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return fmt.Errorf("load config env error,%s", err)
	}
	return loadGlobal()
}

// 加载全局实例
func loadGlobal() (err error) {
	// 加载db全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return fmt.Errorf("load db conn error,%s", err)
	}
	return
}
