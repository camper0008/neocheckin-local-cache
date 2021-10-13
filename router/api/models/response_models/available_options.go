package response_models

import "neocheckin_cache/router/api/models/exported_models"

type AvailableOptions struct {
	Options []exported_models.Option `json:"options"`
}
