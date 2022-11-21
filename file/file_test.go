package file

import (
	"testing"
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
	err := LoadIniFile(&iniConfig, "./file/config.ini")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", iniConfig)
}

func TestFile_LoadJsonFile(t *testing.T) {
	t.Log("TestFile_LoadJsonFile")
	var jsonConfig Config
	err := LoadJsonFile(&jsonConfig, "./file/config.json")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", jsonConfig)
}

func TestFile_LoadYamlFile(t *testing.T) {
	t.Log("TestFile_LoadYamlFile")
	var yamlConfig Config
	err := LoadYamlFile(&yamlConfig, "./file/config.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", yamlConfig)
}
