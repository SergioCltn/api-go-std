package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/sergiocltn/api-go-std/controllers"
	repositories "github.com/sergiocltn/api-go-std/repository"
	"github.com/sergiocltn/api-go-std/routes"
	"github.com/sergiocltn/api-go-std/services"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "root", "secret", "api_go_sstd")

	userRepository := repositories.NewUserRepositoryDB(connStr)

	userService := services.NewUserService(userRepository)
	userController := controllers.NewController(userService)
	router := routes.NewRouter(*userController)

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      router.SetupRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Could not start server", "error", err)
			os.Exit(1)
		}
	}()

	<-done
	fmt.Println("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Failed", "error", err)
		os.Exit(1)
	}
	fmt.Println("Server Exited Properly")
}
