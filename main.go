package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/mkadiri/golang-microservice/healthcheck"
	"github.com/mkadiri/golang-microservice/router"
	"github.com/mkadiri/golang-microservice/usr"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("--- Starting server")
	corsHandler := getCorsHandler()
	initRouter()

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, corsHandler(router.Router)),
		Addr:         fmt.Sprintf(":%d", 8000),
	}

	fmt.Println("--- Listening")
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func getCorsHandler() func(http.Handler) http.Handler {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	return handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-type"}),
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
	)
}

func initRouter()  {
	da
	router.Init()
	usr.InitRouter()
	healthcheck.InitRouter()
}