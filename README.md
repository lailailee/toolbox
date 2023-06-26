## file

### LoadFile 加载ini,json,yaml文件

`func LoadFile(config interface{}, file string, fileType int) (err error)`

```golang

package main

import (
	"github.com/lailailee/toolbox/file"
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

var config Config

file.LoadFile(&config, "config.ini", file.Ini)
file.LoadFile(&config, "config.json", file.Json)
file.LoadFile(&config, "config.yaml", file.Yaml)
```

  

