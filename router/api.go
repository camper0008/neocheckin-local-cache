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
		Handler: api.PostEmployeeCardscan,
	})
	router.Register(Endpoint{
		Path:    `/employee/[^\\]+`,
		Method:  "GET",
		Handler: api.GetEmployeeFromRfid,
	})

	return router
}
