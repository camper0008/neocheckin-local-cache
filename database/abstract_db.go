package database

import (
	"neocheckin_cache/database/database_models"
	"neocheckin_cache/shared"
)

type AbstractDatabase interface {
	GetEmployeeWithRfid(string) database_models.Employee
	GetEmployeeWithDatabaseId(string) database_models.Employee
	InsertEmployee(database_models.Employee)
	UpdateEmployeeWithDatabaseId(string, database_models.Employee)
	DeleteEmployeeWithDatabaseId(string, database_models.Employee)

	GetOptionWithWrapperId(shared.WrapperEnum) database_models.Option
	GetOptionWithDatabaseId(string) database_models.Option
	InsertOption(database_models.Option)
	UpdateOptionWithDatabaseId(string, database_models.Option)
	DeleteOptionWithDatabaseId(string, database_models.Option)

	AddAction(database_models.Action)
}
