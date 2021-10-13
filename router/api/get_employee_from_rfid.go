package api

import (
	"fmt"
	db "neocheckin_cache/database"
	em "neocheckin_cache/router/api/models/exported_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
)

func getRfidFromPath(p string) string {
	a := make([]byte, len(p))
	for c := range p {
		a = append(a, p[c])
		if p[c] == '/' {
			a = make([]byte, len(p)-c)
		}
	}

	return string(a)
}

func GetEmployeeFromRfid(rw http.ResponseWriter, rq http.Request, db db.AbstractDatabase) {
	p := rq.URL.Path
	rfid := getRfidFromPath(p)

	empl, err := db.GetEmployeeWithRfid(rfid)

	if err != nil {
		encoded, err := utils.JsonEncode(rsm.Error{
			Error: err.Error(),
		})

		if err == nil {
			fmt.Fprintf(rw, "%s", encoded)
			return
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
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
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
