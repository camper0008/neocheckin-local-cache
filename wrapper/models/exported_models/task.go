package exported_models

import (
	"time"
)

type Task struct {
	TaskId    int       `json:"taskId"`
	Name      int       `json:"name"`
	Rfid      string    `json:"employeeId"`
	PostKey   string    `json:"highLevelApiKey"`
	SystemId  string    `json:"systemIdentifier"`
	Timestamp time.Time `json:"timestamp"`
}
