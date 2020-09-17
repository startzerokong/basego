package util

import (
	"github.com/mitchellh/mapstructure"
	"github.com/startzerokong/basego/define"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func getRedisPath() ([]byte, error) {
	return getPath("/config/redis.yml")
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

	err := yaml.Unmarshal(configFile, &config)

	if err != nil {

	}

	retConfig := define.Redis{}
	err = mapstructure.Decode(config.Product, &retConfig)

	return &retConfig
}
