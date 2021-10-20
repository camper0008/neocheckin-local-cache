package main

import (
	"fmt"
	"neocheckin_cache/database"
	m "neocheckin_cache/database/models"
	"neocheckin_cache/router"
	"net/http"
)

func main() {
	db := database.MemoryDatabase{}
	db.InsertEmployee(m.Employee{
		Rfid:       "rfid",
		Name:       "rfid",
		Flex:       50,
		Working:    true,
		Department: "rfid",
		Photo:      "rfid",
	})
	db.InsertOption(m.Option{
		WrapperId: 0,
		Name:      "default",
		Schedule: m.Schedule{
			From: m.ScheduleTime{
				Second: 0,
				Minute: 0,
				Hour:   0,
			},
			To: m.ScheduleTime{
				Second: 59,
				Minute: 59,
				Hour:   23,
			},
			Days: m.ScheduleDays{
				Monday:    true,
				Tuesday:   true,
				Wednesday: true,
				Thursday:  true,
				Friday:    true,
				Saturday:  true,
				Sunday:    true,
			},
		},
	})
	router := router.ConnectAPI(&db)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		router.Handle(rw, *r, &db)
	})
	err := http.ListenAndServe(":6000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
