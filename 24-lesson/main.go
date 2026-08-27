package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	port := "8084"
	http.HandleFunc("/", myHandler)
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
