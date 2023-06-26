package toolbox

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// SaveBase64ToImage 将base64图片保存到指定路径
func SaveBase64ToImage(picture string, finalPath string) (err error) {
	// 将没有base64头的picture从base64保存为图片
	if !strings.HasPrefix(picture, "data:image/jpeg;base64,") {
		picture = "data:image/jpeg;base64," + picture
	}
	// 将图片解码
	data, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(picture, "data:image/jpeg;base64,"))
	if err != nil {
		return fmt.Errorf("decode base64 image error:%v", err)
	}
	// 将图片保存到finalPath
	f, err := os.Create(finalPath)
	if err != nil {
		return fmt.Errorf("create file error:%v", err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("write file error:%v", err)
	}
	return nil
}

// ParseImageToBase64 是通过读取图片，并把图片转为base64格式，方便通过json传输
func ParseImageToBase64(url string) (data string, e error) {
	if image, e0 := ioutil.ReadFile(url); e0 != nil {
		e = fmt.Errorf("open file error[%v]", e0)
	} else {
		data = base64.StdEncoding.EncodeToString(image)
	}
	return
}
