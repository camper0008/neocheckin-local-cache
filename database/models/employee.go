package models

type Employee struct {
	DatabaseId string
	WrapperId  int
	Rfid       string
	Name       string
	Flex       int
	Working    bool
	Department string
	Photo      string
}
