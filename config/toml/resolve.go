// Copyright 2017 <614422099@qq.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// toml file resolve.
//
// https://github.com/dingdayu/golangtools
package toml

import (
	"github.com/BurntSushi/toml"
)

// 参考文章 http://www.cnblogs.com/ghj1976/p/4082323.html

func Unmarshal(f []byte, result interface{}) error {
	_, err := toml.Decode(string(f), result)
	return err
}
