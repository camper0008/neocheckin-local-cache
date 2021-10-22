package response_models

import im "neocheckin_cache/wrapper/models/imported_models"

type GetTaskTypes struct {
	Data []im.TaskType `json:"data"`
}
