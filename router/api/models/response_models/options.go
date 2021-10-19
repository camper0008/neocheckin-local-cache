package response_models

import em "neocheckin_cache/router/api/models/exported_models"

type AvailableOptions struct {
	Options []em.Option `json:"options"`
}
