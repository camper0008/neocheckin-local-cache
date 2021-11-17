package config

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func Read() map[string]string {
	parsed := map[string]string{}

	content, err := ioutil.ReadFile("config/settings.conf")
	if err != nil {
		log.Fatal("Unable to read 'config/settings.conf':", err)
	}
	ParseContent(content, parsed)

	return parsed
}

func ParseContent(content []byte, r map[string]string) {
	lines := strings.Split(string(content), "\n")
	for i := 0; i < len(lines); i++ {
		ParseConfLineWithRegex(lines[i], r)
	}
}

func ParseConfLineWithRegex(line string, r map[string]string) {
	exp := regexp.MustCompile(`^([^\s]+)=([^\s#]+)`)
	if exp.MatchString(line) {
		matches := exp.FindStringSubmatch(line)
		k, v := matches[1], matches[2]
		r[k] = v
	}
}
