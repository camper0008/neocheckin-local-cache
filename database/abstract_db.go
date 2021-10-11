package database

import (
	models "neocheckin_cache/database/database_models"
)

type AbstractDatabase interface {
	AddAction(models.Action)
	GetEmployeeFromRfid(string) models.Employee
	UpdateWorkingStatus(models.Employee, bool, int)
	GetWorkingEmployees() models.WorkingEmployees
	GetAvailableOptions() models.Option
}
