package main

import (
	"fmt"
	"log"
	"neocheckin_cache/database"
	"neocheckin_cache/router"
	w "neocheckin_cache/wrapper"
	"net/http"
)

func synchronizeWrapperAndCache(db database.AbstractDatabase) {

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
}

func main() {
	db := database.MemoryDatabase{}

	fmt.Println("Attempting to synchronize...")
	synchronizeWrapperAndCache(&db)
	fmt.Println("Done")

	router := router.ConnectAPI(&db)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		router.Handle(rw, *r, &db)
	})
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
