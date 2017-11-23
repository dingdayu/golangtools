package json

import (
	"encoding/json"
	"bytes"
)

func Unmarshal(f []byte, result interface{}) error {

	// 不推荐使用json.Unmarshal
	decoder := json.NewDecoder(bytes.NewReader(f))
	decoder.UseNumber() // 此处能够保证bigint的精度
	return decoder.Decode(result)
}
