package main

import (
	"log"
	"neocheckin_cache/database"
	"neocheckin_cache/router"
	"neocheckin_cache/utils"
	w "neocheckin_cache/wrapper"
	"net/http"
)

func setupApiServer(db database.AbstractDatabase, l *utils.Logger) {
	router := router.ConnectAPI(db)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		router.Handle(rw, *r, db, l)
	})
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	logger := utils.Logger{}
	logger.CreateLogFile()

	db := database.MemoryDatabase{}
	w.SyncWrapperAndCache(&db, &logger)
	setupApiServer(&db, &logger)
}
