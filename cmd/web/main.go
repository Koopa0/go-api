package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

type Config struct {
}

func main() {

	app := Config{}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
