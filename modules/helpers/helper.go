package helpers

import (
	"io"
	"os"
)

// WriteToFile 往文件里写入内容
func WriteToFile(filename string, content []byte) error {
	//打开文件写入
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	//写入
	_, err = io.WriteString(file, string(content))
	if err != nil {
		return err
	}
	return nil
}
