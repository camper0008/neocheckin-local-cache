package models

type Task struct {
	DatabaseModel
	TaskId       int
	Name         string
	EmployeeRfid string
	SystemId     string
	Timestamp    string // ISO
}
