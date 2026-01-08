package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KimNattanan/exprec-backend/pkg/httpserver"
)

func Start() {
	cfg, db, err := setupDependencies("development")
	if err != nil {
		log.Fatalf("failed to setup dependencies: %v", err)
	}
	app := setupRestServer(db, cfg)

	httpserver.Start(app, cfg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")

	httpserver.Shutdown(app)

	log.Println("Server shutted down.")
}
