package function

import (
	"io/ioutil"
)

// FileGetContents 把整个文件读入一个字符串中
func FileGetContents(filename string) (str string, err error) {
	var data []byte
	data, err = ioutil.ReadFile(filename)
	str = string(data)
	return
}

// FilePutContents 把一个字符串写入文件中
func FilePutContents(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}
