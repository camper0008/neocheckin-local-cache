package utils_test

// FIXME: test rest of logs code

import (
	"io/ioutil"
	"neocheckin_cache/utils"
	"os"
	"strings"
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
		e := "D2015-03-07_T11-06-39"
		o := utils.GetFormattedDate(i)
		if e != o {
			t.Errorf("expected %q, got %q", e, o)
		}
	}
	{
		i, _ := time.Parse(time.UnixDate, "Mon Mar 9 11:36:20 PST 2012")
		e := "D2012-03-09_T11-36-20"
		o := utils.GetFormattedDate(i)
		if e != o {
			t.Errorf("expected %q, got %q", e, o)
		}
	}
}

func TestCreateLogFile(t *testing.T) {
	{
		l := utils.Logger{}
		err := l.CreateLogFile()
		if err != nil {
			t.Error("should not error")
		}
		os.RemoveAll("logs")
	}
	{
		l := utils.Logger{}
		l.CreateLogFile()
		if l.Filename == "" {
			t.Error("should change filename")
		}
		os.RemoveAll("logs")
	}
	{
		l := utils.Logger{}
		l.CreateLogFile()
		_, err := os.Stat(l.Filename)
		if os.IsNotExist(err) {
			t.Error("file should exist")
		}
		os.RemoveAll("logs")
	}
}
func TestAppendToLogFile(t *testing.T) {
	{
		i := "abcdefghijklmn"
		l := utils.Logger{}
		l.CreateLogFile()
		l.AppendToLogFile(i)
		c, _ := ioutil.ReadFile(l.Filename)
		if !strings.Contains(string(c), i) {
			t.Error("file should contain input string after appending")
		}

		os.RemoveAll("logs")
	}
	{
		l := utils.Logger{}
		err := l.AppendToLogFile("")
		if err == nil {
			t.Error("should error with blank filename")
		}
	}
}
func TestFormatAndAppendToLogFile(t *testing.T) {
	{
		i := "abcdefghijklmn"
		l := utils.Logger{}
		l.CreateLogFile()
		l.FormatAndAppendToLogFile(i)
		c, _ := ioutil.ReadFile(l.Filename)
		if !strings.Contains(string(c), i) {
			t.Error("file should contain input string after appending")
		}

		os.RemoveAll("logs")
	}
	{
		l := utils.Logger{}
		err := l.FormatAndAppendToLogFile("")
		if err == nil {
			t.Error("should error with blank filename")
		}
	}
}
