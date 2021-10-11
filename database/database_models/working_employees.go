package database_models

type WorkingEmployees struct {
	Employees []Employee
	Ordered   map[string][]Employee
}
