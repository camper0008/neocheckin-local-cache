package models

type Task struct {
	DatabaseId   string
	TaskId       int
	Name         string
	EmployeeRfid string
	SystemId     string
	Timestamp    string // ISO
}
