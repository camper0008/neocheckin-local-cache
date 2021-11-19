package api

import (
	"fmt"
	dbt "neocheckin_cache/database"
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

	e := []em.Employee{}
	o := map[string][]em.Employee{}

	for i := range dbE {
		if dbE[i].Working {
			genE := em.Employee{
				Name:       dbE[i].Name,
				Flex:       dbE[i].Flex,
				Working:    dbE[i].Working,
				Department: dbE[i].Department,
				Photo:      dbE[i].Photo,
			}
			e = append(e, genE)
			if o[genE.Department] == nil {
				o[genE.Department] = []em.Employee{}
			}
			o[genE.Department] = append(o[genE.Department], genE)
		}
	}

	return rsm.WorkingEmployees{
		Employees: e,
		Ordered:   o,
	}, nil
}
