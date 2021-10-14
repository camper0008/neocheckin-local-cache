package api

import (
	"fmt"
	db "neocheckin_cache/database"
	m "neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	rsm "neocheckin_cache/router/api/models/response_models"
	"neocheckin_cache/utils"
	"net/http"
)

func getRfidFromPath(p string) string {
	i := 0
	a := make([]byte, len(p))
	for c := range p {
		a[c-i] = p[c]
		if p[c] == '/' {
			i = c + 1
			a = make([]byte, len(p)-i)
		}
	}

	return string(a)
}

func GetEmployeeFromRfidEndpoint(rw http.ResponseWriter, rq http.Request, db db.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

	p := rq.URL.Path
	rfid := getRfidFromPath(p)

	empl, err := db.GetEmployeeWithRfid(rfid)

	if err == nil {
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

func GetEmployeeFromRfid(db db.AbstractDatabase, rfid string) (m.Employee, error) {
	empl, err := db.GetEmployeeWithRfid(rfid)
	if err != nil {
		return m.Employee{}, fmt.Errorf("not found")
	}
	return empl, nil
}
