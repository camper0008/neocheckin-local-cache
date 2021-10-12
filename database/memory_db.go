package database

import (
	"fmt"
	"neocheckin_cache/database/database_models"
	"neocheckin_cache/shared"
)

type MemoryDatabase struct {
	AbstractDatabase
	employees []database_models.Employee
	options   []database_models.Option
}

func find[T any](a []T, f func(T) bool) T, error {
	for _, v := range a {
		if f(v) {
			return v
		}
	}
	panic("Not found")
}

func (db *MemoryDatabase) GetEmployeeWithRfid(rfid string) (database_models.Employee, error) {
	for i := range db.employees {
		if db.employees[i].Rfid == rfid {
			return db.employees[i], nil
		}
	}
	panic(fmt.Sprintf("Could not find Employee with rfid '%s'", rfid))
}

func (db *MemoryDatabase) GetEmployeeWithDatabaseId(string) database_models.Employee {
	panic("Not implemented")
}

func (db *MemoryDatabase) InsertEmployee(database_models.Employee) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) UpdateEmployeeWithDatabaseId(string, database_models.Employee) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) DeleteEmployeeWithDatabaseId(string, database_models.Employee) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) GetOptionWithWrapperId(shared.WrapperEnum) (database_models.Option, error) {
	panic("Not implemented")
}

func (db *MemoryDatabase) GetOptionWithDatabaseId(string) (database_models.Option, error) {
	panic("Not implemented")
}

func (db *MemoryDatabase) InsertOption(database_models.Option) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) UpdateOptionWithDatabaseId(string, database_models.Option) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) DeleteOptionWithDatabaseId(string, database_models.Option) error {
	panic("Not implemented")
}

func (db *MemoryDatabase) AddAction(database_models.Action) {
	panic("Not implemented")
}
