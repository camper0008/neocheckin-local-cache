package database_models

type Employee struct {
	DatabaseModel
	Rfid       string
	Name       string
	Flex       int
	Working    bool
	Department string
	Photo      string
}
