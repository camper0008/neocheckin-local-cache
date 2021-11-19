package main

import (
	"fmt"
	"log"
	"neocheckin_cache/database"
	"neocheckin_cache/router"
	"neocheckin_cache/utils"
	w "neocheckin_cache/wrapper"
	"net/http"
)

// FIXME untested
func synchronizeWrapperAndCache(db database.AbstractDatabase, l *utils.Logger) {

	fmt.Println("Attempting to synchronize...")
	t, err := w.GetTaskTypes(l)
	if err != nil {
		fmt.Printf("Error synchronizing task types: %v\n", err)
		l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing task types: %v", err))
	} else {
		w.UpdateDbFromTaskTypes(db, t)
	}

	e, err := w.GetEmployees(l)
	if err != nil {
		fmt.Printf("Error synchronizing employees: %v\n", err)
		l.FormatAndAppendToLogFile(fmt.Sprintf("Error synchronizing employees: %v", err))
	} else {
		w.UpdateDbFromEmployees(db, e)
	}
	fmt.Println("Done")
}

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
	synchronizeWrapperAndCache(&db, &logger)
	setupApiServer(&db, &logger)
}
