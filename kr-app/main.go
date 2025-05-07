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

	r.HandleFunc("/krs", handlers.GetAllKRs).Methods("GET")
	r.HandleFunc("/krs", handlers.CreateKR).Methods("POST")
	r.HandleFunc("/krs/{id}", handlers.GetKR).Methods("GET")
	r.HandleFunc("/krs/{id}", handlers.UpdateKR).Methods("PUT")
	r.HandleFunc("/krs/{id}", handlers.DeleteKR).Methods("DELETE")

	r.HandleFunc("/krs/{kr_id}/subtasks", handlers.GetSubTasks).Methods("GET")
	r.HandleFunc("/krs/{kr_id}/subtasks", handlers.CreateSubTask).Methods("POST")
	r.HandleFunc("/subtasks/{id}", handlers.UpdateSubTask).Methods("PUT")
	r.HandleFunc("/subtasks/{id}", handlers.DeleteSubTask).Methods("DELETE")

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
