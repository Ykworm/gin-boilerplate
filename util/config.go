package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var config *Config

func init() {
	once := sync.Once{}
	once.Do(func() {
		config = &Config{}
	})
}

func GetConfig() *Config {
	return config
}

type Config struct {
	configMap map[string]interface{}
	env       string
}

func (c *Config) GetEnv() string {
	return c.env
}

func (c *Config) InitConfig(env string) {
	c.env = env
	configByte, err := os.ReadFile(fmt.Sprintf("./configs/config-%s.yml", env))
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(configByte, &c.configMap); err != nil {
		panic(err)
	}

}

func (c *Config) GetString(key string) string {
	keys := strings.Split(key, ".")

	if len(keys) == 0 {
		return ""
	} else {
		nextConfigMap := c.configMap
		for i := 0; i < len(keys); i++ {
			if i != len(keys)-1 {
				value := nextConfigMap[keys[i]]
				if value != nil && reflect.TypeOf(value).Kind() == reflect.Map {
					nextConfigMap = value.(map[string]interface{})
				} else {
					nextConfigMap = nil
				}
			}

			if i == len(keys)-1 {
				if nextConfigMap == nil || nextConfigMap[keys[i]] == nil {
					return ""
				}
				return fmt.Sprint(nextConfigMap[keys[i]])
			}
		}
	}
	return ""
}

func (c *Config) GetInt(key string) int {
	value := c.GetString(key)
	i, _ := strconv.Atoi(value)
	return i
}
