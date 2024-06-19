package myrouter

import "net/http"

type Chain struct {
	mw []func(handlerFunc http.HandlerFunc) http.HandlerFunc
}

func NewChain(mw ...func(handlerFunc http.HandlerFunc) http.HandlerFunc) *Chain {
	return &Chain{mw}
}

func (c *Chain) AddMiddleware(mw ...func(http.HandlerFunc) http.HandlerFunc) *Chain {
	c.mw = append(c.mw, mw...)
	return c
}

func (c *Chain) Endpoint(h http.HandlerFunc) http.HandlerFunc {
	for i := len(c.mw) - 1; i >= 0; i-- {
		h = c.mw[i](h)
	}

	return h
}
