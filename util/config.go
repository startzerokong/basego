package util

import (
	"github.com/mitchellh/mapstructure"
	"github.com/startzerokong/basego/define"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func getMemCachePath() ([]byte, error) {
	return getPath("/config/memcache.yml")
}

func getRedisPath() ([]byte, error) {
	return getPath("/config/redis.yml")
}

func getDbPath() ([]byte, error) {
	return getPath("/config/mysql.yml")
}

func getIpFrequencyPath() ([]byte, error) {
	return getPath("/config/ipfrequency.yml")
}

func getUserFrequencyPath() ([]byte, error) {
	return getPath("/config/userfrequency.yml")
}

func getApiPath() ([]byte, error) {
	return getPath("/config/api.yml")
}

func getLogPath() ([]byte, error) {
	return getPath("/config/log.yml")
}

func getPath(path string) ([]byte,error) {
	nowPath ,_ := os.Getwd()
	realPath := nowPath + path

	_, err := os.OpenFile(realPath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(realPath)
}

func GetMemCacheConfig() *define.MemCache {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getMemCachePath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.MemCache{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetRedisConfig() *define.Redis {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getRedisPath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.Redis{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetDbConfig() *define.Db {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getDbPath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.Db{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetIpFrequencyConfig() *define.Ip {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getIpFrequencyPath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.Ip{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetUserFrequencyConfig() *define.User {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getUserFrequencyPath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.User{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetApiConfig() *define.Api {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getApiPath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.Api{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetLogConfig() *define.Log {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	configFile, _ := getLogPath()

	var config envConfig

	_ = yaml.Unmarshal(configFile, &config)

	retConfig := define.Log{}
	_ = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}

func GetConfig(configType string) interface{} {
	type envConfig struct {
		Product  interface{} `yaml:"product"`
		Preline  interface{} `yaml:"preline"`
		Develop  interface{} `yaml:"develop"`
		Test     interface{} `yaml:"test"`
		UnitTest interface{} `yaml:"unit_test"`
	}

	var configFile []byte
	var retConfig interface{}

	if configType == "db" {
		configFile, _ = getDbPath()
		retConfig = define.Db{}
	} else if configType == "redis" {
		configFile, _ = getRedisPath()
		retConfig = define.Redis{}
	} else if configType == "memcache" {
		configFile, _ = getMemCachePath()
		retConfig = define.MemCache{}
	}

	var config envConfig

	err := yaml.Unmarshal(configFile, &config)
	if err != nil {

	}

	err = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}
