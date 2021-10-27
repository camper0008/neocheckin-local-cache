package config

import (
	"io/ioutil"
	"log"
	"strings"
)

func Read() map[string]string {
	r := map[string]string{}

	content, err := ioutil.ReadFile("config/settings.conf")
	if err != nil {
		log.Fatal("Unable to read 'config/settings.conf':", err)
	}
	lines := strings.Split(string(content), "\n")
	for i := 0; i < len(lines); i++ {
		s := strings.SplitN(lines[i], "=", 2)
		k, v := s[0], s[1]
		v = strings.Split(v, "#")[0]
		r[k] = v
	}

	return r
}
