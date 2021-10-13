package exported_models

type Employee struct {
	Name       string `json:"name"`
	Flex       int    `json:"flex"`
	Working    bool   `json:"working"`
	Department string `json:"department"`
	Photo      string `json:"photo"`
}
