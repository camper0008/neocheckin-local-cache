package router

import (
	dbt "neocheckin_cache/database"
	"neocheckin_cache/utils"
	"net/http"
	"regexp"
)

type Endpoint struct {
	Path    string
	Handler func(http.ResponseWriter, http.Request, dbt.AbstractDatabase,
		*utils.Logger)
	Method string
}

type Router struct {
	Path      string
	endpoints []Endpoint
}

func (r *Router) Register(e Endpoint) {
	r.endpoints = append(r.endpoints, e)
}

func (r *Router) Handle(rw http.ResponseWriter, rq http.Request, db dbt.AbstractDatabase, l *utils.Logger) {
	for i := range r.endpoints {
		if endpointMatches(rq.URL.Path, rq.Method, r.Path+r.endpoints[i].Path, r.endpoints[i].Method) {
			r.endpoints[i].Handler(rw, rq, db, l)
		}
	}
}

func endpointMatches(rqPath string, rqMethod string, rtPath string, rtMethod string) bool {
	reP := regexp.MustCompile("^" + rtPath + "$")
	return (reP.FindString(rqPath) != "" && rqMethod == rtMethod)
}
