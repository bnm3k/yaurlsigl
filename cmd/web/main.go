package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/nagamocha3000/yaurlsigl/pkg/store"
	"github.com/pkg/errors"
)

// Config stores configuration variables for the application
type Config struct {
	Addr   string
	DBPath string
}

func getConfig() *Config {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", "localhost:4000", "http network address")
	flag.StringVar(&cfg.DBPath, "dbPath", "./my.db", "path to db file for bolt db")
	return cfg
}

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	store    *store.Store
}

func main() {
	// config
	flag.Parse()
	cfg := getConfig()

	// logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// store
	store, err := store.NewStore(cfg.DBPath)
	if err != nil {
		errorLog.Fatal(errors.Wrapf(err, "Error on store init"))
	}

	// application
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		store:    store,
	}

	// server
	server := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	// start
	app.infoLog.Println("Starting server on", cfg.Addr)
	err = server.ListenAndServe()
	app.errorLog.Fatal(err)
}
