package router

import (
	c "neocheckin_cache/config"
	dbt "neocheckin_cache/database"
	api "neocheckin_cache/router/api"
	"net/http"
)

func HeaderIsValid(h http.Header) bool {
	conf := c.Read()
	token := h.Get("Token")

	return token == conf["CACHE_GET_KEY"]
}

func ConnectAPI(db dbt.AbstractDatabase) Router {
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
	router.Register(Endpoint{
		Path:    "/options",
		Method:  "GET",
		Handler: api.GetOptionsEndpoint,
	})

	return router
}
