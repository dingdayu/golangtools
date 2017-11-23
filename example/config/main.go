package main

import (
	"fmt"
	"github.com/dingdayu/golangtools/config"
	"flag"
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

	file := flag.String("c","conf.json","config file path")
	flag.Parse()

	var conf Config
	err := config.New(*file, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}