package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

func ParseBody(rq http.Request, m interface{}) error {
	headerContentType := rq.Header.Get("Content-Type")
	r := regexp.MustCompile("")
	if r.FindString("application/json") == "" {
		return fmt.Errorf("invalid content type, got '%s', expected 'application/json'", headerContentType)
	}
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(rq.Body)
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
