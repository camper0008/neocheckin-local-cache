package utils

import (
	"io"
	c "neocheckin_cache/config"
	"net/http"
)

// FIXME jeg ved ikke om koden virker
func CreateGetRequest(endpoint string) (*http.Request, error) {
	conf := c.Read()

	req, err := http.NewRequest("GET", conf["WRAPPER_URL"]+endpoint, nil)
	req.Header.Add("token", conf["WRAPPER_KEY"])

	if err != nil {
		return nil, err
	}

	return req, nil
}

// FIXME jeg ved ikke om koden virker
func CreatePostRequest(endpoint string, body io.Reader) (*http.Request, error) {
	conf := c.Read()

	req, err := http.NewRequest("POST", conf["WRAPPER_URL"]+endpoint, body)
	req.Header.Add("token", conf["WRAPPER_KEY"])
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	return req, nil
}
