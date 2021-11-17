package utils_test

// FIXME: test rest of logs code

import (
	"neocheckin_cache/utils"
	"os"
	"testing"
	"time"
)

func TestCreateLogsFolder(t *testing.T) {
	{
		err0 := utils.CreateLogsFolder()
		if err0 != nil {
			t.Error("should not error first time")
		}
		err1 := utils.CreateLogsFolder()
		if err1 == nil {
			t.Error("should error second time")
		}
		f, err2 := os.Create("logs/test.txt")
		if err2 != nil {
			t.Errorf("should exist, got %q", err2.Error())
		}
		f.Close()
		os.RemoveAll("logs")
	}
}
func TestGetFormattedDate(t *testing.T) {
	{
		i, _ := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
		e := "D-2015_03_07-T-11_06_39"
		o := utils.GetFormattedDate(i)
		if e != o {
			t.Errorf("expected %q, got %q", e, o)
		}
	}
	{
		i, _ := time.Parse(time.UnixDate, "Mon Mar 9 11:36:20 PST 2012")
		e := "D-2012_03_09-T-11_36_20"
		o := utils.GetFormattedDate(i)
		if e != o {
			t.Errorf("expected %q, got %q", e, o)
		}
	}
}
