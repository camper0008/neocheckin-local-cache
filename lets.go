package main

import (
	"fmt"
	"log"
	"neocheckin_cache/database"
	"neocheckin_cache/router"
	w "neocheckin_cache/wrapper"
	"net/http"
)

// FIXME jeg ved ikke om funktionen virker
func synchronizeWrapperAndCache(db database.AbstractDatabase) {

	fmt.Println("Attempting to synchronize...")
	t, err := w.GetTaskTypes()
	if err != nil {
		fmt.Printf("Error synchronizing task types: %v\n", err)
	} else {
		w.UpdateDbFromTaskTypes(db, t)
	}

	e, err := w.GetEmployees()
	if err != nil {
		fmt.Printf("Error synchronizing employees: %v\n", err)
	} else {
		w.UpdateDbFromEmployees(db, e)
	}
	fmt.Println("Done")
}

func setupApiServer(db database.AbstractDatabase) {
	router := router.ConnectAPI(db)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		router.Handle(rw, *r, db)
	})
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	db := database.MemoryDatabase{}
	synchronizeWrapperAndCache(&db)
	setupApiServer(&db)
}
