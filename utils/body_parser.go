package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type ParseableBody struct {
	Body   io.ReadCloser
	Header http.Header
}

// FIXME test failer med 0.0% coverage ❌❌ DOG TEST 💪💪💪
// FIXME extract funktioner
func ParseBody(pb ParseableBody, m interface{}) error {
	headerContentType := pb.Header.Get("Content-Type")
	r := regexp.MustCompile("application/json")
	if r.FindString(headerContentType) == "" {
		return fmt.Errorf("invalid content type, got '%s', expected 'application/json'", headerContentType)
	}
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(pb.Body)
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
