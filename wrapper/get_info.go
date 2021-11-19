package wrapper

import (
	"fmt"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
	rm "neocheckin_cache/wrapper/models/response_models"
	"net/http"
)

// FIXME jeg ved ikke om koden virker
func GetTaskTypes(l *utils.Logger) (rm.GetTaskTypes, error) {
	req, err := utils.CreateGetRequest("/tasks/types")
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred creating request: %q", err.Error()))
		return rm.GetTaskTypes{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred doing request: %q", err.Error()))
		return rm.GetTaskTypes{}, err
	}

	parsed := rm.GetTaskTypes{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred parsing request body: %q", err.Error()))
		return rm.GetTaskTypes{}, err
	}

	return parsed, nil
}

// FIXME jeg ved ikke om koden virker
func GetEmployees(l *utils.Logger) (rm.GetEmployees, error) {
	req, err := utils.CreateGetRequest("/employees/all")
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred creating request: %q", err.Error()))
		return rm.GetEmployees{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred doing request: %q", err.Error()))
		return rm.GetEmployees{}, err
	}

	parsed := rm.GetEmployees{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred parsing request body: %q", err.Error()))
		return rm.GetEmployees{}, err
	}

	return parsed, nil
}

// FIXME jeg ved ikke om koden virker
func taskTypesResponseToDbModels(r rm.GetTaskTypes) []dbm.Option {
	data := r.Data
	res := make([]dbm.Option, len(data))
	for i := 0; i < len(data); i++ {
		// FIXME mange linjer, ingen information
		res[i] = dbm.Option{
			WrapperId:   data[i].Id,
			Name:        data[i].Name,
			DisplayName: data[i].DisplayName,
			Priority:    data[i].Priority,
			Locations:   data[i].Locations,
			Category:    data[i].Category,
			Schedule: dbm.Schedule{
				From: dbm.ScheduleTime{
					Second: data[i].Schedule.From.Second,
					Minute: data[i].Schedule.From.Minute,
					Hour:   data[i].Schedule.From.Hour,
				},
				To: dbm.ScheduleTime{
					Second: data[i].Schedule.To.Second,
					Minute: data[i].Schedule.To.Minute,
					Hour:   data[i].Schedule.To.Hour,
				},
				Days: dbm.ScheduleDays{
					Monday:    data[i].Schedule.Days.Monday,
					Tuesday:   data[i].Schedule.Days.Tuesday,
					Wednesday: data[i].Schedule.Days.Wednesday,
					Thursday:  data[i].Schedule.Days.Thursday,
					Friday:    data[i].Schedule.Days.Friday,
					Saturday:  data[i].Schedule.Days.Saturday,
					Sunday:    data[i].Schedule.Days.Sunday,
				},
			},
		}
	}
	return res
}

// FIXME jeg ved ikke om koden virker
func employeesResponseToDbModels(r rm.GetEmployees) []dbm.Employee {
	data := r.Data
	res := make([]dbm.Employee, len(data))
	for i := 0; i < len(data); i++ {
		// FIXME mange linjer, ingen information
		res[i] = dbm.Employee{
			WrapperId:  data[i].WrapperId,
			Rfid:       data[i].Rfid,
			Name:       data[i].Name,
			Flex:       data[i].Flex,
			Working:    data[i].Working,
			Department: data[i].Department,
			Photo:      data[i].Photo,
		}
	}
	return res
}

// FIXME jeg ved ikke om koden virker
func UpdateDbFromTaskTypes(db dbt.AbstractDatabase, r rm.GetTaskTypes) error {
	o := taskTypesResponseToDbModels(r)
	err := db.ReplaceOptions(o)
	return err
}

// FIXME jeg ved ikke om koden virker
func UpdateDbFromEmployees(db dbt.AbstractDatabase, r rm.GetEmployees) error {
	e := employeesResponseToDbModels(r)
	err := db.ReplaceEmployees(e)
	return err
}
