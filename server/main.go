/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aggregator/routing"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func config() {
	viper.AutomaticEnv()
	necessaireVariables := []string{
		"JSERVER1_PORT",
		"JSERVER1_NAME",
		"JSERVER2_PORT",
		"JSERVER2_NAME",
		"SERVER_PORT",
	}

	for i := 0; i < len(necessaireVariables); i++ {
		if !viper.IsSet(necessaireVariables[i]) {
			log.Fatalf("ERREUR : La variable d'environnement '%s' est manquante.", necessaireVariables[i])
		}

	}
}

func main() {
	config()
	fmt.Println("starting point !")

	router := routing.Router()

	http.ListenAndServe(":3008", router)
}
