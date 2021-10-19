package response_models

import em "neocheckin_cache/router/api/models/exported_models"

type Options struct {
	Options []em.Option `json:"options"`
}
