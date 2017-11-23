package yaml

import (
	"github.com/goinggo/mapstructure"
	"gopkg.in/yaml.v2"
)

func Unmarshal(f []byte, result interface{}) error {
	var conf interface{}
	yaml.Unmarshal(f, &conf)
	return mapstructure.Decode(conf, result)
}
