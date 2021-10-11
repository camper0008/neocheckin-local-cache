package main

import (
	r "neocheckin_cache/router"
	"net/http"
)

func main() {
	r.ConnectAPI()
	http.ListenAndServe(":8079", nil)
}
