package response_models

import (
	"neocheckin_cache/router/api/models/exported_models"
)

type WorkingEmployees struct {
	Employees []exported_models.Employee            `json:"employees"`
	Ordered   map[string][]exported_models.Employee `json:"ordered"`
}
