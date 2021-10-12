package response_models

import "neocheckin_cache/router/api/models/exported_models"

type GetEmployee struct {
	Employee []exported_models.Employee
}
