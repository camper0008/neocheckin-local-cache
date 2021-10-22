package main

import (
	"fmt"
	"log"
	"neocheckin_cache/database"
	"neocheckin_cache/router"
	w "neocheckin_cache/wrapper"
	"net/http"
)

func synchronizeWrapperAndCache(db database.AbstractDatabase) error {

	t, err := w.GetTaskTypes()
	if err != nil {
		return err
	} else {
		w.UpdateDbFromTaskTypes(db, t)
	}

	e, err := w.GetEmployees()
	if err != nil {
		return err
	} else {
		w.UpdateDbFromEmployees(db, e)
	}

	return nil

}

func main() {
	db := database.MemoryDatabase{}

	syncErr := synchronizeWrapperAndCache(&db)
	if syncErr != nil {
		fmt.Printf("error occured synchronizing: %q\n", syncErr.Error())
	}

	router := router.ConnectAPI(&db)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		router.Handle(rw, *r, &db)
	})
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
