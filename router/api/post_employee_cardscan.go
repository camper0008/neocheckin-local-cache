package api

import (
	"fmt"
	dbt "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	rqm "neocheckin_cache/router/api/models/request_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	wr "neocheckin_cache/wrapper"
	wem "neocheckin_cache/wrapper/models/exported_models"
	"net/http"
)

func validatePostEmployeeCardscanEndpointInput(rw http.ResponseWriter, p rqm.CardScanned) error {
	missing := []string{}
	if p.ApiKey == "" {
		missing = append(missing, "apiKey")
	}
	if p.EmployeeRfid == "" {
		missing = append(missing, "employeeRfid")
	}
	if p.SystemId == "" {
		missing = append(missing, "systemId")
	}
	if p.Timestamp == "" {
		missing = append(missing, "timestamp")
	}
	if len(missing) == 0 {
		return nil
	} else {
		return fmt.Errorf("missing fields: %v", missing)
	}
}

func PostEmployeeCardscanEndpoint(rw http.ResponseWriter, rq http.Request, db dbt.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

	var p rqm.CardScanned
	err := utils.ParseBody(utils.ParseableBody{Body: rq.Body, Header: rq.Header}, &p)

	if err != nil {
		utils.WriteError(rw, err)
		return
	}

	if err := validatePostEmployeeCardscanEndpointInput(rw, p); err != nil {
		utils.WriteError(rw, err)
		return
	}

	empl, err := db.GetEmployeeWithRfid(p.EmployeeRfid)

	if err == nil {
		// TODO: fix bandaid, use category instead of hardcoding 0 for v0.2

		if p.Option == 0 {
			empl.Working = true
		} else {
			empl.Working = false
		}

		db.UpdateEmployeeWithDatabaseId(empl.DatabaseId, dbm.Employee{
			Rfid:       empl.Rfid,
			Name:       empl.Name,
			Flex:       empl.Flex,
			Working:    empl.Working,
			Department: empl.Department,
			Photo:      empl.Photo,
		})

		statusCode, err := wr.SendTask(wem.Task{
			TaskId:       p.Option,
			Name:         "Scan Card",
			EmployeeRfid: p.EmployeeRfid,
			PostKey:      p.ApiKey,
			SystemId:     p.SystemId,
			Timestamp:    p.Timestamp,
		}, db, false)

		if err != nil {
			if statusCode == http.StatusBadRequest {
				utils.WriteError(rw, err)
			} else if statusCode == http.StatusInternalServerError {
				utils.WriteServerError(rw, err)
			}
			return
		}

		encoded, err := utils.JsonEncode(rsm.GetEmployee{
			Employee: em.Employee{
				Name:       empl.Name,
				Flex:       empl.Flex,
				Working:    empl.Working,
				Department: empl.Department,
				Photo:      empl.Photo,
			},
		})

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
