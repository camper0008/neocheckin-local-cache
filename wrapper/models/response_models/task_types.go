package response_models

type TaskType struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Active      bool     `json:"active"`
	Schedule    Schedule `json:"schedule"`
}

type ScheduleTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}
type ScheduleDays struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
}
type Schedule struct {
	From ScheduleTime
	To   ScheduleTime
	Days ScheduleDays
}

type GetTaskTypes struct {
	Data []TaskType `json:"data"`
}
