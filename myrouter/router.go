package myrouter

import "net/http"

type Router struct {
	http.ServeMux
}

func (r *Router) Get(path string, h http.HandlerFunc) {
	path = http.MethodGet + " " + path
	r.ServeMux.HandleFunc(path, h)
}

func (r *Router) Post(path string, h http.HandlerFunc) {
	path = http.MethodPost + " " + path
	r.ServeMux.HandleFunc(path, h)
}

func (r *Router) AddRouterGroup(g *RouterGroup) {
	for _, route := range g.routes {
		h := g.mwChain.Endpoint(route.Handler)
		r.ServeMux.HandleFunc(route.Method+" "+g.Prefix+route.Path, h)
	}
	for _, sg := range g.subgroups {
		r.addSubgroup(g.Prefix, sg)
	}
}

func (r *Router) addSubgroup(outerPrefix string, sg *RouterGroup) {
	sg.Prefix = outerPrefix + sg.Prefix
	r.AddRouterGroup(sg)
}
