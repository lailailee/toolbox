package toolbox

import (
	"fmt"
	"net"
	"time"
)

// CheckTcpPort 检查端口是否开放
func CheckTcpPort(host string, port string) (e error) {
	remote := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.DialTimeout("tcp", remote, 3*time.Second) // 查看是否连接成功
	if err != nil {
		e = fmt.Errorf("[%v]", err)
		return
	}
	conn.Close()
	return
}
