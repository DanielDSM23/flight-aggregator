/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting point !")

	server := http.NewServeMux()

	server.HandleFunc("GET /health", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(responseWriter, "got path\n")
	})
	server.HandleFunc("GET /flights", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(responseWriter, "handling flights request")
	})
	http.ListenAndServe("localhost:3008", server)
}
