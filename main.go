package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"my-auth-api/handlers"
	"my-auth-api/redis"
)

func main() {
	// Conectar a Redis
	redis.ConnectRedis()

	// Crear router
	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
