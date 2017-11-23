# Config Example

Golang Tools Example Config

## Step A

#### 1、install dep

```
go get -u github.com/golang/dep/cmd/dep
```

> [Golang Dep](https://github.com/golang/dep)

#### 2、 install vendor

```
dep ensure
```

## Step B

```
go get github.com/dingdayu/golangtools/example/config
```

## RUN

```
config -c example/config/conf.json
```

```
{8080 [{gitbook /gitbook [ls ll] dingdayu}]}
```