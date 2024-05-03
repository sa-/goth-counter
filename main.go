package main

import (
	"context"
	"fmt"
	"net/http"

	"goth/templates"
)

var counter = 0

func main() {
	// Setup routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.CounterPage(counter)
		component.Render(context.Background(), w)
	})

	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		counter++
		w.Write([]byte(fmt.Sprintf("%d", counter)))
	})

	// Static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Start server
	port := ":42069"
	fmt.Println("http://localhost" + port)
	http.ListenAndServe(port, nil)
}
