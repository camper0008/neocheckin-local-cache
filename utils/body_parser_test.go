package utils_test

import (
	"io"
	"neocheckin_cache/utils"
	"net/http"
	"strings"
	"testing"
)

type exampleRequest struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	WrongType int    `json:"wrongType"`
	unexposed string
}

func TestBodyParser(t *testing.T) {
	{
		r := http.Request{
			Header: http.Header{
				http.CanonicalHeaderKey("Content-Type"): []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader("Hello, world!")),
		}
		b := exampleRequest{}
		err := utils.ParseBody(utils.ParseableBody{Body: r.Body, Header: r.Header}, &b)

		if !(err != nil && strings.Contains(err.Error(), "invalid character 'H'")) {
			t.Error("Should return JSON error")
		}
	}
	{
		r := http.Request{
			Header: http.Header{
				http.CanonicalHeaderKey("Content-Type"): []string{"text/plain"},
			},
			Body: io.NopCloser(strings.NewReader("{ name: \"Soelberg\", age: 50 }")),
		}
		b := exampleRequest{}
		err := utils.ParseBody(utils.ParseableBody{Body: r.Body, Header: r.Header}, &b)

		if !(err != nil && strings.Contains(err.Error(), "invalid content type")) {
			t.Error("Should deny because of invalid header")
		}
	}
	{
		r := http.Request{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader("{ \"name\": \"Soelberg\", \"age\": 50}")),
		}
		b := exampleRequest{}
		err := utils.ParseBody(utils.ParseableBody{Body: r.Body, Header: r.Header}, &b)

		if err != nil {
			t.Errorf("Should parse correctly, got %q", err.Error())
		}
		if b.Name != "Soelberg" {
			t.Errorf("Should parse name correctly, expected %q got %q", "Soelberg", b.Name)
		}
		if b.Age != 50 {
			t.Errorf("Should parse age correctly, expected %d got %d", 50, b.Age)
		}
	}
	{
		r := http.Request{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader("{ \"name\": \"Soelberg\", \"age\": 50, \"unexposed\": \"none\"}")),
		}
		b := exampleRequest{}
		err := utils.ParseBody(utils.ParseableBody{Body: r.Body, Header: r.Header}, &b)

		if err == nil {
			t.Error("Should error")
		}
		if !(err != nil && strings.Contains(err.Error(), "json: unknown field \"unexposed\"")) {
			t.Error("Should error because of invalid field")
		}
		if b.unexposed != "" {
			t.Errorf("Unexposed should be blank, got '%s'", b.unexposed)
		}
	}
	{
		r := http.Request{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: io.NopCloser(strings.NewReader("{ \"name\": \"Soelberg\", \"age\": 50, \"WrongType\": \"none\"}")),
		}
		b := exampleRequest{}
		err := utils.ParseBody(utils.ParseableBody{Body: r.Body, Header: r.Header}, &b)

		if err == nil {
			t.Error("Should error")
		}
		if !(err != nil && strings.Contains(err.Error(), "Wrong Type provided for field")) {
			t.Error("Should error because of invalid field")
		}
	}
}
