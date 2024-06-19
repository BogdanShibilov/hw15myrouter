package myrouter

import "net/http"

type route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type RouterGroup struct {
	mwChain   Chain
	routes    []route
	Prefix    string
	subgroups []*RouterGroup
}

func NewGroup(prefix string) *RouterGroup {
	return &RouterGroup{
		Prefix: prefix,
	}
}

func (g *RouterGroup) AddSubgroup(sg *RouterGroup) {
	g.subgroups = append(g.subgroups, sg)
}

func (g *RouterGroup) AddMiddleware(mw ...func(http.HandlerFunc) http.HandlerFunc) {
	g.mwChain.AddMiddleware(mw...)
}

func (g *RouterGroup) Get(path string, handler http.HandlerFunc) {
	g.routes = append(g.routes, route{Method: http.MethodGet, Path: path, Handler: handler})
}

func (g *RouterGroup) Post(path string, handler http.HandlerFunc) {
	g.routes = append(g.routes, route{Method: http.MethodPost, Path: path, Handler: handler})
}
