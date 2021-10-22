package request_models

type CardScanned struct {
	EmployeeRfid string `json:"employeeRfid"`
	Option       int    `json:"option"`
	Timestamp    string `json:"timestamp"`
	ApiKey       string `json:"apiKey"`
	SystemId     string `json:"systemId"`
}
