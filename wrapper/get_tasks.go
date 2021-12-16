package wrapper

import (
	"fmt"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	"neocheckin_cache/utils"
	im "neocheckin_cache/wrapper/models/imported_models"
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

	defer resp.Body.Close()

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
func taskTypesResponseToDbModels(r rm.GetTaskTypes) []dbm.Option {
	data := r.Data
	res := make([]dbm.Option, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = taskTypeToDatabaseOption(data[i])
	}
	return res
}

// FIXME jeg ved ikke om koden virker
func updateDbFromTaskTypesResponse(db dbt.AbstractDatabase, r rm.GetTaskTypes) error {
	o := taskTypesResponseToDbModels(r)
	err := db.ReplaceOptions(o)
	return err
}

func taskTypeToDatabaseOption(t im.TaskType) dbm.Option {
	return dbm.Option{
		WrapperId:   t.Id,
		Name:        t.Name,
		DisplayName: t.DisplayName,
		Priority:    t.Priority,
		Locations:   t.Locations,
		Category:    t.Category,
		Schedule: dbm.Schedule{
			From: dbm.ScheduleTime{
				Second: t.Schedule.From.Second,
				Minute: t.Schedule.From.Minute,
				Hour:   t.Schedule.From.Hour,
			},
			To: dbm.ScheduleTime{
				Second: t.Schedule.To.Second,
				Minute: t.Schedule.To.Minute,
				Hour:   t.Schedule.To.Hour,
			},
			Days: dbm.ScheduleDays{
				Monday:    t.Schedule.Days.Monday,
				Tuesday:   t.Schedule.Days.Tuesday,
				Wednesday: t.Schedule.Days.Wednesday,
				Thursday:  t.Schedule.Days.Thursday,
				Friday:    t.Schedule.Days.Friday,
				Saturday:  t.Schedule.Days.Saturday,
				Sunday:    t.Schedule.Days.Sunday,
			},
		},
	}
}
