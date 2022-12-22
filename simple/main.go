/*
To define a simple web server, the minium requirment is to implement the http.Handler interface from net/http package.

which has the following signature: 

package http
type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request)
}

*/


package main

import (
	"fmt"
	"log"
	"net/http"
)

type courses map[string]int

func (c courses) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(rw, "hello golang!")
	case "/list":
		for course, price := range c {
			fmt.Fprintf(rw, "course: %s, price: %d\n", course, price)
		}
	case "/price":
		query := r.URL.Query().Get("course")
		p, ok := c[query]
		if !ok {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(rw, "Course is not available!")
			return 
		}
		fmt.Fprintf(rw, "price: %d", p)
	default:
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "invalid url!")
	}
}

func main() {
	c := courses{
		"golang-beginner": 1000,
		"golang-intermediate": 5000,
		"golang-advanced": 10000,
	}
	log.Fatal(http.ListenAndServe("localhost:8000", c))
}