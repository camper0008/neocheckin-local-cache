package database

import (
	"neocheckin_cache/database/models"
	"neocheckin_cache/shared"
)

type AbstractDatabase interface {
	GetEmployeeWithRfid(string) (models.Employee, error)
	GetEmployeeWithDatabaseId(string) models.Employee
	InsertEmployee(models.Employee) error
	UpdateEmployeeWithDatabaseId(string, models.Employee) error
	DeleteEmployeeWithDatabaseId(string, models.Employee) error

	GetOptionWithWrapperId(shared.WrapperEnum) (models.Option, error)
	GetOptionWithDatabaseId(string) (models.Option, error)
	InsertOption(models.Option) error
	UpdateOptionWithDatabaseId(string, models.Option) error
	DeleteOptionWithDatabaseId(string, models.Option) error

	AddAction(models.Action) error
}
