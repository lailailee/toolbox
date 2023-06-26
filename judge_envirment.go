package toolbox

import (
	"fmt"
	"runtime"
	// "vcapture/core"
)

type Env string

const (
	Darwin  Env = "darwin"
	Linux   Env = "linux"
	Windows Env = "windows"
)

var Envirment = Darwin

// InitEnvirment 初始化环境
func InitEnvirment(env Env) {
	Envirment = env
}

// IsDev 判断当前是否是开发环境 true 开发环境
func IsDev() bool {
	return runtime.GOOS == fmt.Sprintf("%v", Envirment)
}
