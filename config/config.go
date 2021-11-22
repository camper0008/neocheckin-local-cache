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
	exp := confParserRegex()
	if exp.MatchString(line) {
		k, v := "", ""
		matches := exp.FindStringSubmatch(line)
		if matches[1] != "" {
			k = matches[1]
			v = getValueFromMatches(matches)
		}
		if k != "" && v != "" {
			r[k] = v
		}
	}
}

func getValueFromMatches(matches []string) string {
	if matches[2] != "" {
		// double quotes
		return matches[2]
	} else if matches[3] != "" {
		// single quotes
		return matches[3]
	} else if matches[4] != "" {
		// no quotes
		return matches[4]
	}
	return ""
}

func confParserRegex() *regexp.Regexp {
	// ^[\s]*			: make sure that whitespace in front of string isnt matched as part of key
	// ([^\s#]+)		: match the key which isnt whitespace or # (comment)
	// =				: an equal sign
	// (?:				: non-capturing group, to match with the OR correctly
	// "([^"]+)"		: capture anything that isnt a " character, that's surrounded by quotes
	// |				: "OR" - try above first OR if unmatched, instead match below
	// '([^']+)'		: capture anything that isnt a ' character, that's surrounded by quotes
	// |				: "OR" - try above first OR if unmatched, instead match below
	// ([^\s#]+)		: capture anything that isnt whitespace or a # (comment)
	// )				: end non-capturing group
	exp := regexp.MustCompile(`^[\s]*([^\s#]+)=(?:"([^"]+)"|'([^']+)'|([^\s#]+))`)
	return exp
}
