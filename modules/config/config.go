package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/zhuchunshu/sf-file/modules/helpers"
	"gopkg.in/yaml.v3"
	"os"
)

var Config = struct {
	APPName    string `default:"SForum File" env:"APPName"`
	Port       uint   `default:"3000" env:"Port"`
	Url        string `required:"true" env:"Url"`
	Token      string `required:"true" env:"Token"`
	UploadPath string `required:"true" env:"UploadPath"`
}{}

func Get(basePath string) {
	// 判断config.yml是否存在
	filename := basePath + "/config.yml"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println("failed creating file: ", err)
		}

		// 写入默认配置
		err = helpers.WriteToFile(filename, GetDefault())
		if err != nil {
			fmt.Println("failed writing file: ", err)
		}

		fmt.Println("config.yml 配置文件已创建，请修改配置后重启程序！")
		os.Exit(0)
	}

	err := configor.Load(&Config, filename)
	if err != nil {
		return
	}
}

// GetDefault 获取默认配置
func GetDefault() []byte {
	config := Config
	config.APPName = "SForum File"
	config.Port = 3000
	config.Url = "http://localhost:3000/uploads"
	config.Token = helpers.RandString(32)
	config.UploadPath = "./uploads"
	out, err := yaml.Marshal(&config)
	if err != nil {
		return nil
	}
	return out
}
