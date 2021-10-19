package api

import (
	"fmt"
	db "neocheckin_cache/database"
	"neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	"neocheckin_cache/router/api/models/request_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
	"time"
)

func PostEmployeeCardscanEndpoint(rw http.ResponseWriter, rq http.Request, db db.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

	var parsed request_models.CardScanned
	utils.ParseBody(rq, &parsed)

	empl, err := db.GetEmployeeWithRfid(parsed.EmployeeRfid)

	if err == nil {
		err := db.AddAction(models.Action{
			Timestamp: time.Now(),
			Option:    parsed.Option,
			Rfid:      empl.Rfid,
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
