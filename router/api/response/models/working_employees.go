package response_models

import (
	"neocheckin_cache/router/api/response/models/exported_models"
)

type WorkingEmployees struct {
	Employees []exported_models.Employee
	Ordered   map[string][]exported_models.Employee
}
