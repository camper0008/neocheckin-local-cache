package imported_models

type Employee struct {
	WrapperId  int    `json:"id"`
	Rfid       string `json:"rfid"`
	Name       string `json:"name"`
	Flex       int    `json:"flex"`
	Working    bool   `json:"working"`
	Department string `json:"department"`
	Photo      string `json:"photo"`
	Error      string `json:"error"`
}
