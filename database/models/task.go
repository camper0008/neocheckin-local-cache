package models

import (
	"time"
)

type Task struct {
	DatabaseModel
	Name      string
	TaskId    int
	Rfid      string
	Option    int
	Timestamp time.Time
}
