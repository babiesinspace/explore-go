package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Hinge/backend-hw-junior/api"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "backend-hw-junior.db")
	if err != nil {
		log.Fatal("Could not connect to sqlite", err)
	}
	server := api.NewServer(db)

	r := chi.NewRouter()
	r.Get("/", server.StatusHandler)
	r.Get("/recommendations/{userID}", server.RecommendationsHandler)

	fmt.Println("Your app is now serving at 127.0.0.1:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
