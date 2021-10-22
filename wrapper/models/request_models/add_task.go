package request_models

type AddTask struct {
	TaskId       int    `json:"taskId"`
	Name         string `json:"name"`
	EmployeeRfid string `json:"employeeRfid"`
	PostKey      string `json:"highLevelApiKey"`
	SystemId     string `json:"systemIdentifier"`
	Timestamp    string `json:"timestamp"` // ISO
}
