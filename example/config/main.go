package main

import (
	"fmt"
	"os/signal"
	"os"
	"github.com/dingdayu/golangtools/config"
)

// 实例
type Instance struct {
	Name 	string
	Path	string
	Cmd		[]string
	User	string
}

// 配置
type Config struct {
	Port int
	Instance []Instance
}

func main() {

	fmt.Println("JSON 配置文件解析：")
	ParserJson();

	fmt.Println("Toml 配置文件解析：")
	ParserToml()

	fmt.Println("Yaml 配置文件解析：")
	ParserYaml()

	fmt.Println("Xml 配置文件解析：")
	ParserXml()



	go func() {
		// Handle interrupt signal
		ch := make(chan os.Signal)
		signal.Notify(ch, os.Interrupt)

		<-ch
	}()
}

func ParserJson()  {
	var conf Config
	err := config.New("conf.json", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}

func ParserYaml()  {
	var conf Config
	err := config.New("conf.yaml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}

func ParserXml()  {
	var conf Config
	err := config.New("conf.xml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}

func ParserToml()  {
	var conf Config
	err := config.New("conf.toml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}