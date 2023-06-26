package test

import (
	"testing"

	"github.com/lailailee/toolbox"
)

type Config struct {
	Owner struct {
		Name         string `ini:"name"`
		Organization string `ini:"organization"`
	} `ini:"owner"`
	Database struct {
		Server   string `ini:"server"`
		Port     int    `ini:"port"`
		Type     string `ini:"type"`
		Username string `ini:"username"`
		Password string `ini:"password"`
	} `ini:"database"`
}

func TestFile_LoadIniFile(t *testing.T) {
	t.Log("TestFile_LoadIniFile")
	var iniConfig Config
	err := toolbox.LoadFile(&iniConfig, "./config.ini", toolbox.Ini)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", iniConfig)
}

func TestFile_LoadJsonFile(t *testing.T) {
	t.Log("TestFile_LoadJsonFile")
	var jsonConfig Config
	err := toolbox.LoadFile(&jsonConfig, "./config.json", toolbox.Json)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", jsonConfig)
}

func TestFile_LoadYamlFile(t *testing.T) {
	t.Log("TestFile_LoadYamlFile")
	var yamlConfig Config
	err := toolbox.LoadFile(&yamlConfig, "./config.yaml", toolbox.Yaml)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", yamlConfig)
}
