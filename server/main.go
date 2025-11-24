/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aggregator/routing"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting point !")

	router := routing.Router()

	http.ListenAndServe(":3008", router)
}
