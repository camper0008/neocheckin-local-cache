package main

import (
	"neocheckin_cache/database"
	"neocheckin_cache/database/models"
	"neocheckin_cache/router"
	"net/http"
)

func main() {
	db := database.MemoryDatabase{}
	db.InsertEmployee(models.Employee{
		Rfid:       "rfid",
		Name:       "rfid",
		Flex:       50,
		Working:    true,
		Department: "rfid",
		Photo:      "rfid",
	})
	router := router.ConnectAPI(&db)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		router.Handle(rw, *r, &db)
	})
	http.ListenAndServe(":8079", nil)
}
