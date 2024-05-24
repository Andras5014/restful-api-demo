package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 配置映射成config对象

func LoadConfigFormToml(filePath string) (*Config, error) {
	config = NewDefaultConfig()
	// 读取配置文件
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return nil, fmt.Errorf("load config file error,path: %s,%s", err, filePath)
	}
	return config, nil
}
func LoadConfigFormEnv() (*Config, error) {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	// 配置文件映射成config对象
	return config, nil
}
