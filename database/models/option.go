package models

type Option struct {
	DatabaseId  string
	WrapperId   int
	Name        string
	DisplayName string
	Locations   []string
	Category    string
	Priority    bool
	Schedule    Schedule
}

type Schedule struct {
	From ScheduleTime
	To   ScheduleTime
	Days ScheduleDays
}
type ScheduleTime struct {
	Hour   int
	Minute int
	Second int
}
type ScheduleDays struct {
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool
	Sunday    bool
}
