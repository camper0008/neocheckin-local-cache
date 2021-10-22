package api

import (
	"fmt"
	dbt "neocheckin_cache/database"
	em "neocheckin_cache/router/api/models/exported_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
)

func GetEmployeesWorkingEndpoint(rw http.ResponseWriter, rq http.Request, db dbt.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

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

func GetEmployeesWorking(db dbt.AbstractDatabase) (rsm.WorkingEmployees, error) {
	dbE, err := db.GetAllEmployees()
	if err != nil {
		return rsm.WorkingEmployees{}, err
	}

	e := make([]em.Employee, len(dbE))
	o := make(map[string][]em.Employee)

	for i := range e {
		e[i] = em.Employee{
			Name:       dbE[i].Name,
			Flex:       dbE[i].Flex,
			Working:    dbE[i].Working,
			Department: dbE[i].Department,
			Photo:      dbE[i].Photo,
		}
		if o[e[i].Department] == nil {
			o[e[i].Department] = []em.Employee{}
		}
		o[e[i].Department] = append(o[e[i].Department], e[i])
	}

	return rsm.WorkingEmployees{
		Employees: e,
		Ordered:   o,
	}, nil
}
