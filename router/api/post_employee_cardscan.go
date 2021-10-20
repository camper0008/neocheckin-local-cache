package api

import (
	"fmt"
	db "neocheckin_cache/database"
	em "neocheckin_cache/router/api/models/exported_models"
	rqm "neocheckin_cache/router/api/models/request_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	wr "neocheckin_cache/wrapper"
	wem "neocheckin_cache/wrapper/models/exported_models"
	"net/http"
)

func validatePostEmployeeCardscanEndpointInput(rw http.ResponseWriter, p rqm.CardScanned) error {
	missing := ""
	if p.ApiKey == "" {
		missing += " apiKey"
	}
	if p.EmployeeRfid == "" {
		missing += " employeeRfid"
	}
	if p.SystemId == "" {
		missing += " systemId"
	}
	if p.Timestamp == "" {
		missing += " timestamp"
	}
	if missing == "" {
		return nil
	} else {
		return fmt.Errorf("missing fields:%s", missing)
	}
}

func PostEmployeeCardscanEndpoint(rw http.ResponseWriter, rq http.Request, db db.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

	var p rqm.CardScanned
	err := utils.ParseBody(rq, &p)

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
		err := wr.SendTask(wem.Task{
			TaskId:    p.Option,
			Name:      "Scan card",
			Rfid:      p.EmployeeRfid,
			PostKey:   p.ApiKey,
			SystemId:  p.SystemId,
			Timestamp: p.Timestamp,
		})

		if err != nil {
			utils.WriteServerError(rw, err)
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
