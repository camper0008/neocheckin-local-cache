package router

import "neocheckin_cache/database"

func ConnectAPI(db database.AbstractDatabase) Router {
	router := Router{}
	return router
}
