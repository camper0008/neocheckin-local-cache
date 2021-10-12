package main

import (
	"neocheckin_cache/router"
	"net/http"
)

func main() {
	router.ConnectAPI()
	http.ListenAndServe(":8079", nil)
}
