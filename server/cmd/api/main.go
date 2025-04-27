package main

import (
	"fmt"
	"log"
	"net/http"
	"testgolangapp/server/internal/database"
	"testgolangapp/server/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := mux.NewRouter()
	
	r.HandleFunc("/api/tasks", handlers.GetAllTasksHandler).Methods("GET")
	r.HandleFunc("/api/tasks/{id:[0-9]+}", handlers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/api/tasks", handlers.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/api/tasks/{id:[0-9]+}", handlers.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/api/tasks/{id:[0-9]+}", handlers.DeleteTaskHandler).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	port := 8000
	fmt.Printf("Server running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
