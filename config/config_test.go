package config_test

import (
	c "neocheckin_cache/config"
	"testing"
)

func TestRead(t *testing.T) {
	{
		p := c.Read()
		if len(p) == 0 {
			t.Error("config should not be blank")
		}
	}
}

func TestParseContent(t *testing.T) {
	{
		content := []byte("a=b\nc=d\ne=f")
		parsed := map[string]string{}
		c.ParseContent(content, parsed)
		if parsed["a"] != "b" || parsed["c"] != "d" || parsed["e"] != "f" {
			t.Error("invalid content")
		}
	}
	{
		content := []byte("a=b\nc=d\ne=f")
		parsed := map[string]string{}
		c.ParseContent(content, parsed)
		if len(parsed) != 3 {
			t.Error("invalid length")
		}
	}
}
