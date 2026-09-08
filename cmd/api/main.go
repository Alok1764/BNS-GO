package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Alok1764/GO/internal/config"
	"github.com/Alok1764/GO/internal/db"
	"github.com/Alok1764/GO/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DB_URL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	lh := handlers.NewListingHandler(db)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server is listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)

	}

}
