package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!"));
	})
	
	fmt.Println("Server up and running!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Failed to start the server!")
	}
}
