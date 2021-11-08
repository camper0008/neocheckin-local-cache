package utils

import (
	"encoding/json"
	"fmt"
)

// FIXME jeg ved ikke om koden virker
func JsonEncode(s interface{}) ([]byte, error) {
	b, err := json.Marshal(s)

	if err != nil {
		return b, fmt.Errorf("could not encode to json: '%s'", err.Error())
	}

	return b, nil
}
