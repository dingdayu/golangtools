// Copyright 2017 <614422099@qq.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// yaml file resolve.
//
// https://github.com/dingdayu/golangtools
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
