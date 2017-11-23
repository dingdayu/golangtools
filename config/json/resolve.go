// Copyright 2017 <614422099@qq.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// json file resolve.
//
// https://github.com/dingdayu/golangtools
package json

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(f []byte, result interface{}) error {

	// 不推荐使用json.Unmarshal
	decoder := json.NewDecoder(bytes.NewReader(f))
	decoder.UseNumber() // 此处能够保证bigint的精度
	return decoder.Decode(result)
}
