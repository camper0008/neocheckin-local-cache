package request_models

type CardScanned struct {
	EmployeeRfid string `json:"employeeRfid"`
	Name         string `json:"name"`
	Option       int    `json:"option"`
	Timestamp    string `json:"timestamp"`
	ApiKey       string `json:"apiKey"`
	SystemId     string `json:"systemId"`
}
