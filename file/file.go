package file

import (
	"encoding/json"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// LoadIniFile 读取ini配置文件
func LoadIniFile(config interface{}, file string) (err error) {
	err = ini.MapTo(config, file)
	return
}

// LoadJsonFile 读取json配置文件
func LoadJsonFile(config interface{}, file string) (err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, config)
	return
}

// LoadYamlFile 读取yaml配置文件
func LoadYamlFile(config interface{}, file string) (err error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	// 解析yaml
	err = yaml.Unmarshal(data, config)
	return
}
