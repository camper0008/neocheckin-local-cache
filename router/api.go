package router

import (
	"neocheckin_cache/database"
	api "neocheckin_cache/router/api"
)

func ConnectAPI(db database.AbstractDatabase) Router {
	router := Router{
		Path: "/api",
	}

	router.Register(Endpoint{
		Path:    `/employee/cardscanned`,
		Method:  "POST",
		Handler: api.PostEmployeeCardscanEndpoint,
	})
	router.Register(Endpoint{
		Path:    `/employees/working`,
		Method:  "GET",
		Handler: api.GetEmployeesWorkingEndpoint,
	})
	router.Register(Endpoint{
		Path:    `/employee/[^\\]+`,
		Method:  "GET",
		Handler: api.GetEmployeeFromRfidEndpoint,
	})

	return router
}
