# Golang Tools Config

Golang Tools Config Support：

- Json
- Yaml
- Toml
- Xml

## Install

```
go get github.com/dingdayu/golangtools/config
```

Or

```
go get gitee.com/dingdayu/golangtools/config
```


## Example

### json

```
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

	var conf Config
	err := config.New("conf.json", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

Run:

```
go run ../example/config/main.go -c ../example/config/conf.json
{8080 [{gitbook /gitbook [ll ls] dingdayu}]}
```

> 请注意路径！

> 参见 [example/config/main.go](example/config/main.go)

### xml


```
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

	var conf Config
	err := config.New("conf.xml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

> 参加 [example/config/main.go](example/config/main.go)

### toml


```
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

	var conf Config
	err := config.New("conf.toml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

> 参见 [example/config/main.go](example/config/main.go)

### yaml


```
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

	var conf Config
	err := config.New("conf.yaml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

> 参见 [example/config/main.go](example/config/main.go)