package api

import (
	db "neocheckin_cache/database"
	dbm "neocheckin_cache/database/models"
	em "neocheckin_cache/router/api/models/exported_models"
	"net/http"
)

func GetOptionsEndpoint(rw http.ResponseWriter, rq http.Request, db db.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")

	//encoded, err := utils.JsonEncode(rsm.GetEmployee{
	//	Employee: em.Employee{
	//		Name:       empl.Name,
	//		Flex:       empl.Flex,
	//		Working:    empl.Working,
	//		Department: empl.Department,
	//		Photo:      empl.Photo,
	//	},
	//})
	//
	//if err == nil {
	//	fmt.Fprintf(rw, "%s", encoded)
	//	return
	//} else {
	//	utils.WriteServerError(rw, err)
	//	return
	//}
}

func GetOptions(db db.AbstractDatabase) ([]dbm.Option, error) {
	o, err := db.GetAllOptions()
	if err != nil {
		return o, nil
	}
	return o, nil
}

func ConvertOptionsToExportedModels(o []dbm.Option) []em.Option {
	result := make([]em.Option, len(o))

	return result
}
