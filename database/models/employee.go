package models

type Employee struct {
	DatabaseModel
	WrapperId  int
	Rfid       string
	Name       string
	Flex       int
	Working    bool
	Department string
	Photo      string
}
