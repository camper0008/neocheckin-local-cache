package utils

import (
	"io"
	c "neocheckin_cache/config"
	"net/http"
)

func CreateGetRequest(endpoint string) (*http.Request, error) {
	conf := c.Read()

	req, err := http.NewRequest("GET", conf["WRAPPER_URL"]+endpoint, nil)
	req.Header.Add("token", conf["WRAPPER_GET_KEY"])

	if err != nil {
		return nil, err
	}

	return req, nil
}

func CreatePostRequest(endpoint string, token string, body io.Reader) (*http.Request, error) {
	conf := c.Read()

	req, err := http.NewRequest("POST", conf["WRAPPER_URL"]+endpoint, body)
	req.Header.Add("token", token)

	if err != nil {
		return nil, err
	}

	return req, nil
}
