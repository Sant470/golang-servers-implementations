/*
ServerMux is request multiplexer which solves the problem of associating different urls to different methods.

type HandlerFunc(rw http.ResponseWriter, r *http.Request)

type (f HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	f(rw, r)
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type muxEntry struct {
	pattern string
	handler Handler
}

type ServerMux struct {
	mu map[string]muxEntry
}

func (mux *ServerMux) Handle(pattern string, h Handler) {
	mux.mu[pattern] = Handler
}

func (mux *ServerMux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Path
	h := mux.mu[pattern]
	h.ServeHTTP(rw, r)
}

One thing to note is that Handle accepts pattern and an instance of Handler, which we are able to acheive by converting our custom method to HandlerFunc (which implments the Handler interface).
There are a lot in net/http package, this is just to have a mental model what's actually happening behind the scene.
*/

package main

import (
	"fmt"
	"net/http"
)

type courses map[string]int

func (c courses) list(rw http.ResponseWriter, r *http.Request) {
	for course, price := range c {
		fmt.Fprintf(rw, "course: %s, price: %d\n", course, price)
	}
}

func (c courses) price(rw http.ResponseWriter, r *http.Request) {
	course := r.URL.Query().Get("course")
	price, ok := c[course]
	if !ok {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "course not available!")
		return 
	}
	fmt.Fprintf(rw, "price: %d", price)
}

func main() {
	c := courses{
		"golang-beginner": 1000,
		"golang-intermediate": 5000,
		"golang-advanced": 10000,
	}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(c.list))
	mux.Handle("/price", http.HandlerFunc(c.price))
	http.ListenAndServe("localhost:8000", mux)
}