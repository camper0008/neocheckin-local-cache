package router

import "net/http"

type Endpoint struct {
	Path    string
	Handler func(*http.ResponseWriter, http.Request)
	Method  string
}

type Router struct {
	Path      string
	Endpoints []Endpoint
}

func (r *Router) Register(e Endpoint) {
	r.Endpoints = append(r.Endpoints, e)
}

func (r *Router) Handle(rw *http.ResponseWriter, rq http.Request) {
	for i := range r.Endpoints {
		if rq.URL.Path == (r.Path+r.Endpoints[i].Path) && rq.Method == r.Endpoints[i].Method {
			r.Endpoints[i].Handler(rw, rq)
		}
	}
}
