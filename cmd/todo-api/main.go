package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/the6fallenangel/go-todo/internal/config"
	"github.com/the6fallenangel/go-todo/internal/storage"
)

func main() {
	cfg := config.Load()

	store, err := storage.New(cfg)
	if err != nil {
		log.Fatal("error initializing storage:", err)
	}

	mux := http.NewServeMux()
	registerRoutes(mux, store)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		fmt.Println("listening on", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server error:", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal("graceful shutdown failed:", err)
	}
	fmt.Println("server stopped cleanly")
}
