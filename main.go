package main

import (
	"fmt"
	"hw15myhttp/myrouter"
	"net/http"
)

func main() {
	r := myrouter.NewRouter()

	postIt := myrouter.NewChain(Middleware1, Middleware2).Endpoint(PostIt)
	r.Post("/post", postIt)

	group := myrouter.NewGroup("/outer")
	group.Get("/first", First)
	second := myrouter.NewChain(Middleware1).Endpoint(Second)
	group.Get("/second", second)

	medium1 := myrouter.NewGroup("/medium1")
	medium1.Get("/third", Third)
	medium1.Post("/postit", PostIt)
	medium1.AddMiddleware(Middleware2)

	medium2 := myrouter.NewGroup("/medium2")
	medium2.Get("/first", First)

	group.AddSubgroup(medium1)
	group.AddSubgroup(medium2)

	fourth := myrouter.NewChain(Middleware1).Endpoint(Fourth)
	inner1 := myrouter.NewGroup("/inner1")
	inner1.Get("/fourth", fourth)
	inner1.Get("/fifth", Fifth)
	medium1.AddSubgroup(inner1)

	inner2 := myrouter.NewGroup("/inner2")
	inner2.Get("/first", First)
	medium2.AddSubgroup(inner2)

	r.AddRouterGroup(group)
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	_ = srv.ListenAndServe()
}

func First(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("First endpoint\n"))
	fmt.Println("Reached first endpoint")
}

func Second(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Second endpoint\n"))
	fmt.Println("Reached second endpoint")
}

func Third(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Third endpoint\n"))
	fmt.Println("Reached third endpoint")
}

func Fourth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Fourth endpoint\n"))
	fmt.Println("Reached fourth endpoint")
}

func Fifth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Fifth endpoint\n"))
	fmt.Println("Reached fifth endpoint")
}

func PostIt(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("PostIt endpoint\n"))
	fmt.Println("Reached postIt endpoint")
}

func Middleware1(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entered middleware 1")
		h(w, r)
		fmt.Println("Exiting middleware 1")
	}
}

func Middleware2(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entered middleware 2")
		h(w, r)
		fmt.Println("Exiting middleware 2")
	}
}
