package utils_test

import (
	"neocheckin_cache/utils"
	"testing"
)

type exampleBody struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Pointer interface{} `json:"pointer"`
}

func TestJsonEncode(t *testing.T) {
	{
		p := 20
		e := exampleBody{
			Name:    "scree",
			Age:     10,
			Pointer: &p,
		}
		_, err := utils.JsonEncode(e)
		if err != nil {
			t.Error("should not error with pointers")
		}
	}
	{
		p := "a"
		e := exampleBody{
			Name:    "scree",
			Age:     10,
			Pointer: p,
		}
		_, err := utils.JsonEncode(e)
		if err != nil {
			t.Errorf("should not error, %q", err.Error())
		}
	}
}
