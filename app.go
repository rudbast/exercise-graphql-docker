package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/rudbast/exercise-graphql-docker/config"
	"github.com/rudbast/exercise-graphql-docker/database"
	"github.com/rudbast/exercise-graphql-docker/web"
)

func main() {
	flag.Parse()

	if err := config.Init(); err != nil {
		log.Fatalf("[ERR] Initiate config error: %+v.\n", err)
	}

	cfg := config.Get()
	if cfg.Debug {
		log.Printf("Config: %+v\n", config.Get())
	}

	if err := database.Init("mysql", cfg.Database.String()); err != nil {
		log.Fatalf("[ERR] Initiate database error: %+v.\n", err)
	}
	defer database.Get().Close()

	r := mux.NewRouter()
	web.NewArticleHandler(r)

	http.Handle("/", r)

	// Serve http.
	log.Fatal(http.ListenAndServe(cfg.Server.Host, nil))
}
