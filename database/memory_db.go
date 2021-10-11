package database

import (
	models "neocheckin_cache/database/database_models"
	"sort"
)

type MemoryDatabase struct {
	AbstractDatabase
	employees []models.Employee
	options   []models.Option
}

func (db *MemoryDatabase) GetEmployeeFromRfid(rfid string) (models.Employee, bool) {
	for i := range db.employees {
		if db.employees[i].Rfid == rfid {
			return db.employees[i], true
		}
	}

	return models.Employee{}, false
}

func (db *MemoryDatabase) UpdateWorkingStatus(employee models.Employee, checkingIn bool, optionId int) {
	employee.Working = checkingIn
}

func (db *MemoryDatabase) GetWorkingEmployees() models.WorkingEmployees {
	ordered := map[string][]models.Employee{}

	for i := range db.employees {
		ordered[db.employees[i].Department] = append(ordered[db.employees[i].Department], db.employees[i])
	}

	for d := range ordered {
		sort.Slice(ordered[d], func(i int, j int) bool {
			return ordered[d][i].Name > ordered[d][j].Name
		})
	}

	return models.WorkingEmployees{
		Employees: db.employees,
		Ordered:   ordered,
	}
}

func (db *MemoryDatabase) GetAvailableOptions() []models.Option {
	return db.options
}
