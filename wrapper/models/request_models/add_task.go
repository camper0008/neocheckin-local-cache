package request_models

type AddTask struct {
	TaskId           int    `json:"taskId"`
	Name             string `json:"name"`
	EmployeeRfid     string `json:"employeeId"`
	ApiKey           string `json:"highLevelApiKey"`
	SystemIdentifier string `json:"systemIdentifier"` // unique identifier for the system, ex. "viborg-klient-01", for logging purposes.
}
