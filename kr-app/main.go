package main

import (
	"log"
	"net/http"

	"kr-app/db"
	"kr-app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.Init() // Inicializa conexión a la base de datos

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")

	r.HandleFunc("/krs", handlers.GetAllKRsHandler).Methods("GET")
	r.HandleFunc("/krs", handlers.CreateKRHandler).Methods("POST")
	r.HandleFunc("/krs/{id}", handlers.UpdateKRHandler).Methods("PUT")
	r.HandleFunc("/krs/{id}", handlers.DeleteKRHandler).Methods("DELETE")
	r.HandleFunc("/krs/{id}/edit", handlers.EditKeyResultHandler).Methods("GET")

	r.HandleFunc("/krs/{id}/subtasks", handlers.GetSubTasksHandler).Methods("GET")
	r.HandleFunc("/krs/{id}/subtasks", handlers.CreateSubTaskHandler).Methods("POST")
	//r.HandleFunc("/subtasks/{id}", handlers.UpdateSubTaskHandler).Methods("PUT")
	r.HandleFunc("/subtasks/{id}", handlers.UpdateSubTaskHandler).Methods("PUT")
	r.HandleFunc("/subtasks/{id}", handlers.DeleteSubTaskHandler).Methods("DELETE")
	r.HandleFunc("/subtasks/{id}/edit", handlers.EditSubTaskFormHandler).Methods("GET")

	r.HandleFunc("/verify", handlers.VerifyTicketStatusHandler).Methods("GET")

	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
