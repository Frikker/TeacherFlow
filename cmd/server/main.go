package main

import (
	"chat-service/config"
	"chat-service/pkg/handlers"
	"chat-service/pkg/middleware/db"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	dbConn, err := db.NewConn(ctx, cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	router := handlers.NewHandler(dbConn)
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router.InitRoutes(), // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	ctxWTimeout, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctxWTimeout)
	os.Exit(0)
}
