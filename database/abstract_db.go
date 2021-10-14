package database

import (
	"neocheckin_cache/database/models"
	"neocheckin_cache/shared"
)

type AbstractDatabase interface {
	GetEmployeeWithRfid(string) (models.Employee, error)
	GetEmployeeWithDatabaseId(string) (models.Employee, error)
	GetAllEmployees(string) (models.Employee, error)
	InsertEmployee(models.Employee) error
	UpdateEmployeeWithDatabaseId(string, models.Employee) error
	DeleteEmployeeWithDatabaseId(string) error

	GetOptionWithWrapperId(shared.WrapperEnum) (models.Option, error)
	GetOptionWithDatabaseId(string) (models.Option, error)
	GetAllOptions(string) (models.Option, error)
	InsertOption(models.Option) error
	UpdateOptionWithDatabaseId(string, models.Option) error
	DeleteOptionWithDatabaseId(string) error

	AddAction(models.Action) error
	GetAllAction(models.Action) error
	DeleteActionWithDatabaseId(string, models.Action) error
}
