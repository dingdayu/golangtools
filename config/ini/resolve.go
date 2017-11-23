package ini

import (
	"gopkg.in/gcfg.v1"
)

func Unmarshal(f []byte, result interface{}) error {
	return gcfg.ReadStringInto(result, string(f))
}
