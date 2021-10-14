package database

import (
	m "neocheckin_cache/database/models"
	"neocheckin_cache/shared"
)

type AbstractDatabase interface {
	GetEmployeeWithRfid(string) (m.Employee, error)
	GetEmployeeWithDatabaseId(string) (m.Employee, error)
	GetAllEmployees() ([]m.Employee, error)
	InsertEmployee(m.Employee) error
	UpdateEmployeeWithDatabaseId(string, m.Employee) error
	DeleteEmployeeWithDatabaseId(string) error

	GetOptionWithWrapperId(shared.WrapperEnum) (m.Option, error)
	GetOptionWithDatabaseId(string) (m.Option, error)
	GetAllOptions() ([]m.Option, error)
	InsertOption(m.Option) error
	UpdateOptionWithDatabaseId(string, m.Option) error
	DeleteOptionWithDatabaseId(string) error

	AddAction(m.Action) error
	GetAllActions() ([]m.Action, error)
	DeleteActionWithDatabaseId(string, m.Action) error
}
