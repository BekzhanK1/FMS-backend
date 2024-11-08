package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-gateway/internal/routes"
	"api-gateway/internal/utils"

	"github.com/gorilla/mux"
)

func Run() {
	utils.InitGRPCClient("localhost:5001")

	router := mux.NewRouter()
	routes.Routes(router)

	port := "8080"
	server := &http.Server{
		Addr:   ":" + port,
		Handler: router,
	}

	go gracefulShutdown(server)

	log.Printf("Server(API) is starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server startup failed: %v\n", err)
	}

	log.Println("Server(API) gracefully stopped")
}

func gracefulShutdown(server *http.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v\n", err)
	}
}