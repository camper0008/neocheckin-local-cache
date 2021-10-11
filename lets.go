package main

import (
	br "neocheckin_cache/database"
	r "neocheckin_cache/router"
	"net/http"
)

func main() {
	br.Jhawdudhwu()
	r.ConnectAPI()
	http.ListenAndServe(":8079", nil)
}
