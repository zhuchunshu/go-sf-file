package helpers

import (
	"io"
	"math/rand"
	"os"
	"strings"
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

// RandString 随机生成一个字符串
func RandString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// IsImage 判断上传的内容是图片
func IsImage(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}
