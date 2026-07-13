package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/internal/config"
	"todo/internal/storage"
)

func main() {
	cfg := config.Load()

	store, err := storage.New(cfg)
	if err != nil {
		log.Fatal("error initializing storage:", err)
	}

	mux := http.NewServeMux()
	registerRoutes(mux, store)

	addr := ":8080"
	fmt.Println("listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
