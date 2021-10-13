package utils

import (
	"encoding/json"
	"fmt"
)

func JsonEncode(s interface{}) ([]byte, error) {
	b, err := json.Marshal(s)

	if err != nil {
		panic(fmt.Sprintf("Could not encode to json: '%s'", err.Error()))
	}

	return b, nil
}
