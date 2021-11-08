package imported_models

type TaskType struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
	Priority    bool     `json:"priority"`
	Schedule    Schedule `json:"schedule"`
	Locations   []string `json:"exclusiveLocations"`
	Category    string   `json:"category"`
	BlankField  bool     `json:"active"` // "Active", though unused for our program
	Error       string   `json:"error"`
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
	From ScheduleTime `json:"from"`
	To   ScheduleTime `json:"to"`
	Days ScheduleDays `json:"days"`
}
