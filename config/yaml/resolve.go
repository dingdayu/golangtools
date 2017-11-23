package yaml

import (
	"gopkg.in/yaml.v2"
	"github.com/goinggo/mapstructure"
)

func Unmarshal(f []byte, result interface{}) error {
	var conf interface{}
	yaml.Unmarshal(f, &conf)
	return mapstructure.Decode(conf, result)
}
