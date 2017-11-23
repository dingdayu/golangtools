package xml

import (
	"encoding/xml"
)

func Unmarshal(f []byte, result interface{}) error {
	return xml.Unmarshal(f, result)
}
