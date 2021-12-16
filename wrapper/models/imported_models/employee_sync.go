package imported_models

type EmployeeSync struct {
	WrapperId  int    `json:"id"`
	Rfid       string `json:"rfid"`
	Name       string `json:"name"`
	Flex       int    `json:"flex"`
	Working    bool   `json:"working"`
	Department string `json:"department"`
	Error      string `json:"error"`
}
