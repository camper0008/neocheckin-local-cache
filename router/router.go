package router

import (
	"neocheckin_cache/database"
	"net/http"
	"regexp"
)

type Endpoint struct {
	Path    string
	Handler func(http.ResponseWriter, http.Request, database.AbstractDatabase)
	Method  string
}

type Router struct {
	Path      string
	endpoints []Endpoint
}

func (r *Router) Register(e Endpoint) {
	r.endpoints = append(r.endpoints, e)
}

func (r *Router) Handle(rw http.ResponseWriter, rq http.Request, db database.AbstractDatabase) {
	rw.Header().Add("Content-Type", "application/json")
	for i := range r.endpoints {
		reP := regexp.MustCompile("^" + r.Path + r.endpoints[i].Path + "$")

		if reP.FindString(rq.URL.Path) != "" && rq.Method == r.endpoints[i].Method {
			r.endpoints[i].Handler(rw, rq, db)
		}
	}
}
