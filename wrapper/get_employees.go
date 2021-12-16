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

func getEmployeesWithoutPhoto(l *utils.Logger) (rm.GetEmployeesSync, error) {
	req, err := utils.CreateGetRequest("/employees/sync")
	if err != nil {
		l.FormatAndAppendToLogFile(fmt.Sprintf("error occurred creating request: %q", err.Error()))
		return rm.GetEmployeesSync{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.FormatAndAppendToLogFile(
			fmt.Sprintf(
				"error occurred doing request: %q",
				err.Error(),
			),
		)
		return rm.GetEmployeesSync{}, err
	}

	defer resp.Body.Close()

	parsed := rm.GetEmployeesSync{}
	err = utils.ParseBody(utils.ParseableBody{
		Body:   resp.Body,
		Header: resp.Header,
	}, &parsed)
	if err != nil {
		l.FormatAndAppendToLogFile(
			fmt.Sprintf(
				"error occurred parsing request body: %q",
				err.Error(),
			),
		)
		return rm.GetEmployeesSync{}, err
	}

	return parsed, nil
}

// FIXME jeg ved ikke om koden virker
func getEmployeesWithPhoto(l *utils.Logger) (rm.GetEmployees, error) {
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

	defer resp.Body.Close()

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

func employeeToDatabaseEmployee(e im.Employee) dbm.Employee {
	return dbm.Employee{
		WrapperId:  e.WrapperId,
		Rfid:       e.Rfid,
		Name:       e.Name,
		Flex:       e.Flex,
		Working:    e.Working,
		Department: e.Department,
		Photo:      e.Photo,
	}
}

// FIXME jeg ved ikke om koden virker
func allEmployeesResponseToDbModels(r rm.GetEmployees) []dbm.Employee {
	data := r.Data
	res := make([]dbm.Employee, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = employeeToDatabaseEmployee(data[i])
	}
	return res
}

func getSyncedDatabaseEmployee(sync im.EmployeeSync, db dbt.AbstractDatabase) (dbm.Employee, error) {
	empl, err := db.GetEmployeeWithRfid(sync.Rfid)
	if err != nil {
		return empl, err
	}
	syncedEmpl := dbm.Employee{
		DatabaseId: empl.DatabaseId,
		Photo:      empl.Photo,
		WrapperId:  sync.WrapperId,
		Name:       sync.Name,
		Department: sync.Department,
		Working:    sync.Working,
		Flex:       sync.Flex,
		Rfid:       sync.Rfid,
	}
	return syncedEmpl, nil
}

func syncEmployeesResponseToDbModels(db dbt.AbstractDatabase, r rm.GetEmployeesSync) ([]dbm.Employee, []error) {
	data := r.Data
	res := make([]dbm.Employee, len(data))
	errors := make([]error, len(data))
	for i := 0; i < len(data); i++ {
		syncedEmpl, err := getSyncedDatabaseEmployee(data[i], db)
		if err != nil {
			res[i] = dbm.Employee{}
			errors[i] = err
		} else {
			res[i] = syncedEmpl
			errors[i] = nil
		}
	}

	return res, errors
}

// FIXME jeg ved ikke om koden virker
func updateDbFromSyncEmployeesResponse(db dbt.AbstractDatabase, r rm.GetEmployeesSync) []error {
	syncedEmployees, errors := syncEmployeesResponseToDbModels(db, r)
	for i := 0; i < len(syncedEmployees); i++ {
		if errors[i] == nil {
			db.UpdateEmployeeWithDatabaseId(syncedEmployees[i].DatabaseId, syncedEmployees[i])
		}
	}
	return errors
}

func updateDbFromAllEmployeesResponse(db dbt.AbstractDatabase, r rm.GetEmployees) error {
	e := allEmployeesResponseToDbModels(r)
	err := db.ReplaceEmployees(e)
	return err
}
