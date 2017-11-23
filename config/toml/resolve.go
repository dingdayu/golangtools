package toml

import (
	"github.com/BurntSushi/toml"
)

// 参考文章 http://www.cnblogs.com/ghj1976/p/4082323.html

func Unmarshal(f []byte, result interface{}) error {
	_, err := toml.Decode(string(f), result)
	return err
}
