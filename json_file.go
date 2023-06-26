package toolbox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadJsonFromFile(filename string, p interface{}) (err error) {
	if f, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		d := []byte(f)
		if err := json.Unmarshal(d, p); err != nil {
			return err
		}
	}
	return
}

func SetJsonToFile(filename string, js string) (err error) {
	var f *os.File
	if CheckFileIsExist(filename) {
		// os.Truncate(filename, 0)
		f, err = os.OpenFile(filename, os.O_RDWR, 0666)
		fmt.Println("文件存在")
	} else {
		f, err = os.Create(filename)
		fmt.Println("文件不存在")
	}
	if err != nil {
		return
	}
	defer f.Close()
	// data, _ := json.Marshal(js)
	_, err = f.Write([]byte(js))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return
	}
	return
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
