package toolbox

import (
	"crypto/md5"
	"fmt"
)

// Md5 将字符串转为md5
func EncryptStringToMd5(str string) string {
	md5String := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5String
}
