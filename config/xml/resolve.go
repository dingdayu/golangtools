// Copyright 2017 <614422099@qq.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// xml file resolve.
//
// https://github.com/dingdayu/golangtools
package xml

import (
	"encoding/xml"
)

func Unmarshal(f []byte, result interface{}) error {
	return xml.Unmarshal(f, result)
}
