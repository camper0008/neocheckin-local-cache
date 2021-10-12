package database

import (
	"fmt"
	"neocheckin_cache/database/models"
	"neocheckin_cache/shared"
)

type MemoryDatabase struct {
	AbstractDatabase
	employees []models.Employee
	options   []models.Option
}

func findEmployee(a []models.Employee, f func(models.Employee) bool) (models.Employee, error) {
	for _, v := range a {
		if f(v) {
			return v, nil
		}
	}
	panic("Not found")
}

func (db *MemoryDatabase) GetEmployeeWithRfid(rfid string) (models.Employee, error) {
	employees, err := findEmployee(db.employees, func(model models.Employee) bool {
		return true
	})

	print(employees, err)

	for i := range db.employees {
		if db.employees[i].Rfid == rfid {
			return db.employees[i], nil
		}
	}
	panic(fmt.Sprintf("Could not find Employee with rfid '%s'", rfid))
}

func (db *MemoryDatabase) GetEmployeeWithDatabaseId(string) models.Employee {
	panic("Not implemented")
}

func (db *MemoryDatabase) InsertEmployee(models.Employee) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) UpdateEmployeeWithDatabaseId(string, models.Employee) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) DeleteEmployeeWithDatabaseId(string, models.Employee) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) GetOptionWithWrapperId(shared.WrapperEnum) (models.Option, error) {
	panic("Not implemented")
}

func (db *MemoryDatabase) GetOptionWithDatabaseId(string) (models.Option, error) {
	panic("Not implemented")
}

func (db *MemoryDatabase) InsertOption(models.Option) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) UpdateOptionWithDatabaseId(string, models.Option) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) DeleteOptionWithDatabaseId(string, models.Option) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) AddAction(models.Action) {
	panic("Not implemented")
}
