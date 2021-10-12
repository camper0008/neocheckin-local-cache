package database

import (
	"neocheckin_cache/database/database_models"
	"neocheckin_cache/shared"
)

type AbstractDatabase interface {
	GetEmployeeWithRfid(string) (database_models.Employee, error)
	GetEmployeeWithDatabaseId(string) database_models.Employee
	InsertEmployee(database_models.Employee) error
	UpdateEmployeeWithDatabaseId(string, database_models.Employee) error
	DeleteEmployeeWithDatabaseId(string, database_models.Employee) error

	GetOptionWithWrapperId(shared.WrapperEnum) (database_models.Option, error)
	GetOptionWithDatabaseId(string) (database_models.Option, error)
	InsertOption(database_models.Option) error
	UpdateOptionWithDatabaseId(string, database_models.Option) error
	DeleteOptionWithDatabaseId(string, database_models.Option) error

	AddAction(database_models.Action) error
}
