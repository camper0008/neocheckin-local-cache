package response_models

import "neocheckin_cache/router/api/models/exported_models"

type CardScanned struct {
	Employee exported_models.Employee `json:"employee"`
}
