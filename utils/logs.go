package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Logger struct {
	Filename string
}

func GetFormattedDate(t time.Time) string {
	return fmt.Sprintf(
		"D%0.4d-%0.2d-%0.2d_T%0.2d-%0.2d-%0.2d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}

func FormatLogMessage(msg string) string {
	t := GetFormattedDate(time.Now())
	return t + ":\n\t" + msg
}

func CreateLogsFolder() error {
	fmt.Println("attempting to create logs folder")
	// os.ModePerm == 0777 (unix permissions for creating directories)
	err := os.Mkdir("logs", os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		fmt.Println(err.Error())
	} else if err != nil && errors.Is(err, os.ErrExist) {
		fmt.Println("logs folder already exists")
	} else if err == nil {
		fmt.Println("logs folder created")
	}
	fmt.Println("done attempting to create logs folder")
	return err
}

func (l *Logger) CreateLogFile() {
	err := CreateLogsFolder()
	if err != nil && !errors.Is(err, os.ErrExist) {
		fmt.Printf("an error occured creating logs folder: %v", err)
		return
	}
	name := GetFormattedDate(time.Now())
	f, err := os.Create("logs/" + name + ".txt")
	if err != nil {
		fmt.Printf("an error occured creating log file: %v", err)
		return
	}
	f.Close()
	l.Filename = "logs/" + name + ".txt"
}

func (l *Logger) FormatAndAppendToLogFile(msg string) error {
	formatted := FormatLogMessage(msg)
	err := l.AppendToLogFile(formatted)
	return err
}
func (l *Logger) AppendToLogFile(msg string) error {
	if l.Filename == "" {
		return fmt.Errorf("logger has not been given a filename")
	}
	f, wErr := os.OpenFile(l.Filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if wErr != nil {
		return wErr
	}
	_, wErr = f.WriteString(msg + "\n\n")
	return wErr
}
