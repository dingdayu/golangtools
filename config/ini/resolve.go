package ini

import (
	"gopkg.in/gcfg.v1"
)

/**
解析ini
*/
func Unmarshal(f []byte, result interface{}) error {
	return gcfg.ReadStringInto(result, string(f))
}
