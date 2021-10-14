package response_models

import (
	em "neocheckin_cache/router/api/models/exported_models"
)

type WorkingEmployees struct {
	Employees []em.Employee            `json:"employees"`
	Ordered   map[string][]em.Employee `json:"ordered"`
}
