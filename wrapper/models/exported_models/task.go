package exported_models

type Task struct {
	TaskId    int    `json:"taskId"`
	Name      string `json:"name"`
	Rfid      string `json:"rfid"`
	PostKey   string `json:"highLevelApiKey"`
	SystemId  string `json:"systemIdentifier"`
	Timestamp string `json:"timestamp"` // ISO
}
