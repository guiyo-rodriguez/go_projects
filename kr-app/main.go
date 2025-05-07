package main

import (
	"html/template"
	"log"
	"net/http"

	"kr-app/db"
	"kr-app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.Init() // Inicializa conexión a la base de datos

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/kr_item.html", "templates/subtask_item.html", "templates/kr_list.html"))
		krs, err := db.GetAllKeyResults()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		tmpl.Execute(w, krs)
	})

	r := mux.NewRouter()

	r.HandleFunc("/krs", handlers.GetAllKRsHandler).Methods("GET")
	r.HandleFunc("/krs", handlers.CreateKRHandler).Methods("POST")
	r.HandleFunc("/krs/{id}", handlers.DeleteKRHandler).Methods("DELETE")
	r.HandleFunc("/krs/{id}/subtasks", handlers.CreateSubTaskHandler).Methods("POST")
	r.HandleFunc("/subtasks/{id}", handlers.UpdateSubTaskHandler).Methods("PUT")
	r.HandleFunc("/subtasks/{id}", handlers.DeleteSubTaskHandler).Methods("DELETE")

	/* r.HandleFunc("/krs", handlers.GetAllKRsHandler).Methods("GET")
	r.HandleFunc("/krs", handlers.CreateKRHandler).Methods("POST")
	r.HandleFunc("/krs/{id}", handlers.GetKRHandler).Methods("GET")
	//r.HandleFunc("/krs/{id}", handlers.UpdateKR).Methods("PUT")
	r.HandleFunc("/krs/{id}", handlers.DeleteKRHandler).Methods("DELETE")

	//r.HandleFunc("/krs/{kr_id}/subtasks", handlers.GetSubTasks).Methods("GET")
	r.HandleFunc("/krs/{kr_id}/subtasks", handlers.CreateSubTaskHandler).Methods("POST")
	r.HandleFunc("/subtasks/{id}", handlers.UpdateSubTaskHandler).Methods("PUT")
	r.HandleFunc("/subtasks/{id}", handlers.DeleteSubTaskHandler).Methods("DELETE")
	*/
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
