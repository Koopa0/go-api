package main

import (
	"flag"
	"fmt"
	"github.com/koopa0/go-api/internal/config"
	"github.com/koopa0/go-api/internal/driver"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"

var counts int64
var app config.AppConfig

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {

	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "jogroup", "Database name")
	dbUser := flag.String("dbuser", "koopa", "Database user")
	dbPass := flag.String("dbpass", "", "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "", "Database ssl setting (disable, prefer,require)")

	flag.Parse()

	if *dbName == "" || *dbUser == "" {
		fmt.Println("Missing required flags")
		os.Exit(1)
	}

	// connect to database
	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		*dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	return db, nil
}
