package api

import (
	"fmt"
	dbt "neocheckin_cache/database"
	"neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
)

func GetEmployeesWorkingEndpoint(rw http.ResponseWriter, rq http.Request, db dbt.AbstractDatabase, l *utils.Logger) {
	rw.Header().Add("Content-Type", "application/json; charset=utf-8")

	w, err := GetEmployeesWorking(db)

	if err == nil {
		encoded, err := utils.JsonEncode(w)

		if err == nil {
			fmt.Fprintf(rw, "%s", encoded)
			return
		} else {
			utils.WriteServerError(rw, err)
			return
		}
	} else {
		utils.WriteError(rw, err)
	}
}

// FIXME for lang, jeg ved ikke om det virker
func GetEmployeesWorking(db dbt.AbstractDatabase) (rsm.WorkingEmployees, error) {
	dbE, err := db.GetAllEmployees()
	if err != nil {
		return rsm.WorkingEmployees{}, err
	}

	// TODO: sort alphabetically

	e, o := getUnorderedAndOrderedEmployees(dbE)

	return rsm.WorkingEmployees{
		Employees: e,
		Ordered:   o,
	}, nil
}

func getUnorderedAndOrderedEmployees(dbE []models.Employee) ([]em.Employee, map[string][]em.Employee) {
	e := []em.Employee{}
	o := map[string][]em.Employee{}

	for i := range dbE {
		if dbE[i].Working {
			convE := convertDBModelToExportedModel(dbE[i])
			e = append(e, convE)
			if o[convE.Department] == nil {
				o[convE.Department] = []em.Employee{}
			}
			o[convE.Department] = append(o[convE.Department], convE)
		}
	}
	SortWorkingEmployees(o)
	return e, o
}

func SortWorkingEmployees(o map[string][]em.Employee) {
	for _, v := range o {
		sortEmployees(v)
	}
}

func sortEmployees(v []em.Employee) {
	for i := 0; i < len(v)-1; i++ {
		for j := i; j < len(v)-1; j++ {
			if utils.CompareString(v[j].Name, v[j+1].Name) > 0 {
				temp := v[j]
				v[j] = v[j+1]
				v[j+1] = temp
			}
		}
	}
}

func convertDBModelToExportedModel(e models.Employee) em.Employee {
	genE := em.Employee{
		Name:       e.Name,
		Flex:       e.Flex,
		Working:    e.Working,
		Department: e.Department,
		Photo:      e.Photo,
	}
	return genE
}
