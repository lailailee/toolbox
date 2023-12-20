package test

import (
	"testing"

	"github.com/lailailee/toolbox"
)

func TestFile_InitLogger(t *testing.T) {
	t.Log("TestFile_LoadIniFile")
	var loggerConfig = toolbox.LoggerConfig{
		Filename:   "./test.log",
		MaxSize:    5,
		MaxBackups: 3,
		MaxAge:     3,
		Compress:   false,
		LocalTime:  true,
	}
	logger := toolbox.InitLogger(loggerConfig, toolbox.DebugLevel)
	logger.Infoln("111")
}
