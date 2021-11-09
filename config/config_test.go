package config_test

/*
note:

TestRead will likely yield or fail because of it being "unable to find `config/settings.conf`"
this is because when testing it will look for `config/settings.conf` in the `config` directory,
because that is where the `config_test.go` file runs.
which, errors because ofcourse it cannot find `config/config/settings.conf`
if you want to pass the first test, simply create a file in `config/config/settings.conf`
and add a line like `test=test`

-tpho
*/

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
