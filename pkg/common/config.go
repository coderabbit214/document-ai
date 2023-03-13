package common

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// GlobalObj 全局参数
type GlobalObj struct {
	//监听ip
	Server Server `yaml:"server"`
	OpenAi OpenAi `yaml:"openai"`
	Qdrant Qdrant `yaml:"qdrant"`
}

// Server 服务器参数
type Server struct {
	//监听ip
	Port string `yaml:"port"`
}

type OpenAi struct {
	Key string `yaml:"key"`
}

type Qdrant struct {
	Host           string `yaml:"host"`
	CollectionName string `yaml:"collection-name"`
}

// GlobalObject 全局属性
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("read config/config.yaml err:", err)
		return
	}
	err = yaml.Unmarshal(file, g)
	if err != nil {
		fmt.Println("conf Unmarshal err:", err)
		return
	}
}

func init() {
	GlobalObject = &GlobalObj{}
	//尝试加载配置文件，获取用户自定义配置
	GlobalObject.Reload()
}
