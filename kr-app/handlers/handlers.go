package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"kr-app/db"
	"kr-app/models"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Mostrar todos los KRs
func GetAllKRsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAllKRsHandler")
	krs, err := db.GetAllKeyResults()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "kr_list.html", krs)
}

// Crear un nuevo KR
func CreateKRHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateKRHandler")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kr := models.KeyResult{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	id, err := db.InsertKeyResult(kr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	kr.ID = id

	// Devolver un fragmento HTML que HTMX puede insertar
	templates.ExecuteTemplate(w, "kr_item.html", kr)
}

// Eliminar un KR
func DeleteKRHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteKRHandler")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := db.DeleteKeyResult(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Agregar sub-tarea a un KR
func CreateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateSubTaskHandler")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	krID, _ := strconv.Atoi(mux.Vars(r)["id"])

	st := models.SubTask{
		KRID:  krID,
		Title: r.FormValue("title"),
		Done:  false,
	}

	id, err := db.InsertSubTask(st)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	st.ID = id

	templates.ExecuteTemplate(w, "subtask_item.html", st)
}

// Actualizar sub-tarea (por ejemplo marcar como hecha)
func UpdateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("UpdateSubTaskHandler")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	done := r.FormValue("done") == "true"

	st := models.SubTask{
		ID:    id,
		Title: r.FormValue("title"),
		Done:  done,
	}

	if err := db.UpdateSubTask(st); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updated, _ := db.GetSubTask(id)
	templates.ExecuteTemplate(w, "subtask_item.html", updated)
}

// Eliminar sub-tarea
func DeleteSubTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("DeleteSubTaskHandler")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := db.DeleteSubTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
