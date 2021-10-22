package database

import (
	m "neocheckin_cache/database/models"
)

type AbstractDatabase interface {
	GetEmployeeWithRfid(string) (m.Employee, error)
	GetEmployeeWithDatabaseId(string) (m.Employee, error)
	GetAllEmployees() ([]m.Employee, error)
	InsertEmployee(m.Employee) error
	UpdateEmployeeWithDatabaseId(string, m.Employee) error
	DeleteEmployeeWithDatabaseId(string) error

	GetOptionWithWrapperId(int) (m.Option, error)
	GetOptionWithDatabaseId(string) (m.Option, error)
	GetAllOptions() ([]m.Option, error)
	InsertOption(m.Option) error
	UpdateOptionWithDatabaseId(string, m.Option) error
	DeleteOptionWithDatabaseId(string) error

	AddTask(m.Task) error
	GetAllTasks() ([]m.Task, error)
	DeleteTaskWithDatabaseId(string) error
}
