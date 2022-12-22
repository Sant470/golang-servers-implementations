/*
Using the global instance of ServeMux which is provided by net http package.
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
	http.Handle("/list", http.HandlerFunc(c.list))
	http.Handle("/price", http.HandlerFunc(c.price))
	http.ListenAndServe("localhost:8000", nil)
}