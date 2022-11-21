package file

import (
	"encoding/json"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	Ini = iota + 0
	Json
	Yaml
)

// LoadFile 按类型读取文件
func LoadFile(config interface{}, file string, fileType int) (err error) {
	switch fileType {
	case Ini:
		err = loadIniFile(config, file)
	case Json:
		err = loadJsonFile(config, file)
	case Yaml:
		err = loadYamlFile(config, file)
	}
	return
}

// loadIniFile 读取ini类型文件
func loadIniFile(config interface{}, file string) (err error) {
	err = ini.MapTo(config, file)
	return
}

// loadJsonFile 读取json类型文件
func loadJsonFile(config interface{}, file string) (err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, config)
	return
}

// loadYamlFile 读取yaml类型文件
func loadYamlFile(config interface{}, file string) (err error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	// 解析yaml
	err = yaml.Unmarshal(data, config)
	return
}
