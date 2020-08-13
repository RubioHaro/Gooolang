package main

import (
	"net/http"
)

// Router struct
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) findHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodIsAllowed := r.rules[path][method]
	return handler, methodIsAllowed, exist
}

func newRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	handler, methodIsAllowed, exist := r.findHandler(request.URL.Path, request.Method)
	if !exist {
		// fmt.Fprintf(w, "404. El recurso solicitado no existe!")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !methodIsAllowed {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		handler(w, request)
	}
}
