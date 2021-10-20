package response_models

import em "neocheckin_cache/router/api/models/exported_models"

type CardScanned struct {
	Employee em.Employee `json:"employee"`
}
