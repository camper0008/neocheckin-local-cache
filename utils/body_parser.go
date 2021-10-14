package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ParseBody(r http.Request, m interface{}) error {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		return fmt.Errorf("invalid content type, got '%s', expected 'application/json'", headerContentType)
	}
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&m)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return fmt.Errorf("bad request. Wrong Type provided for field %s" + unmarshalErr.Field)
		} else {
			return fmt.Errorf("bad request %s", err.Error())
		}
	}
	return nil
}
