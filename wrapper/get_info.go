// TODO: log all errors
package wrapper

import (
	c "neocheckin_cache/config"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
	rm "neocheckin_cache/wrapper/models/response_models"
	"net/http"
)

func createRequest(endpoint string) (*http.Request, error) {
	conf := c.Read()

	req, err := http.NewRequest("GET", conf["WRAPPER_URL"]+endpoint, nil)
	req.Header.Add("token", conf["WRAPPER_GET_KEY"])

	if err != nil {
		return nil, err
	}

	return req, nil
}

func GetTaskTypes() (rm.GetTaskTypes, error) {
	req, err := createRequest("/tasks/types")
	if err != nil {
		return rm.GetTaskTypes{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return rm.GetTaskTypes{}, err
	}

	parsed := rm.GetTaskTypes{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		return rm.GetTaskTypes{}, err
	}

	return parsed, nil
}

func GetEmployees() (rm.GetEmployees, error) {
	req, err := createRequest("/employees/all")
	if err != nil {
		return rm.GetEmployees{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return rm.GetEmployees{}, err
	}

	parsed := rm.GetEmployees{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		return rm.GetEmployees{}, err
	}

	return parsed, nil
}

func taskTypesResponseToDbModels(r rm.GetTaskTypes) []dbm.Option {
	data := r.Data
	res := make([]dbm.Option, len(data))
	for i := 0; i < len(data); i++ {
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
func employeesResponseToDbModels(r rm.GetEmployees) []dbm.Employee {
	data := r.Data
	res := make([]dbm.Employee, len(data))
	for i := 0; i < len(data); i++ {
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

func UpdateDbFromTaskTypes(db dbt.AbstractDatabase, r rm.GetTaskTypes) error {
	o := taskTypesResponseToDbModels(r)
	err := db.ReplaceOptions(o)
	return err
}

func UpdateDbFromEmployees(db dbt.AbstractDatabase, r rm.GetEmployees) error {
	e := employeesResponseToDbModels(r)
	err := db.ReplaceEmployees(e)
	return err
}
