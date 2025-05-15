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

	sectors, err := db.GetAllSectors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := models.PageDataKeyResults{
		KeyResults: krs,
		Sectors:    sectors,
	}

	log.Printf("page data: %v\n", data)

	templates.ExecuteTemplate(w, "kr_list.html", data)
}

// Crear un nuevo KR
func CreateKRHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateKRHandler: %v\n", r.Body)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sector_id, _ := strconv.Atoi(r.FormValue("sector_id"))
	kr := models.KeyResult{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		SectorID:    sector_id,
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
	log.Printf("UpdateSubTaskHandler, r: %v\n", r.Body)
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

	log.Printf("UpdateSubTaskHandler, st: %v\n", st)

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

func GetSubTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetSubTaskHandler")
	vars := mux.Vars(r)
	krIDStr := vars["id"]
	krID, err := strconv.Atoi(krIDStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	kr, err := db.GetKeyResult(krID)
	if err != nil {
		http.Error(w, "KR no encontrado", http.StatusNotFound)
		return
	}

	log.Printf("KR: %v", kr)

	//subtasks, err := db.GetSubTasksByKRID(krID)
	subtasks, err := db.GetSubTasksByKRID(krID)
	if err != nil {
		http.Error(w, "Error obteniendo subtareas", http.StatusInternalServerError)
		return
	}

	log.Printf("subtasks: %v", subtasks)

	// tmpl := template.Must(template.ParseFiles(
	// 	"templates/subtasks_view.html",
	// ))

	data := struct {
		KR       models.KeyResult
		SubTasks []models.SubTask
	}{
		KR:       kr,
		SubTasks: subtasks,
	}

	//log.Printf("tmpl: %v", tmpl)
	log.Printf("data: %v", data)

	//tmpl.Execute(w, data)
	templates.ExecuteTemplate(w, "subtasks_view.html", data)
}

func EditSubTaskFormHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("EditSubTaskFormHandler: %v\n", r.Body)
	id, _ := strconv.Atoi(vars["id"])
	st, err := db.GetSubTask(id)
	if err != nil {
		http.Error(w, "Subtarea no encontrada", 404)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/subtask_edit_form.html"))
	tmpl.Execute(w, st)
}

/*
	func UpdateSubTaskHandler(w http.ResponseWriter, r *http.Request) {
		log.Printf("UpdateSubTaskHandler: %v\n", r.Body)
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		title := r.FormValue("title")

		err := db.UpdateSubTaskTitle(id, title)
		if err != nil {
			http.Error(w, "No se pudo actualizar", 500)
			return
		}

		st, _ := db.GetSubTask(id)

		tmpl := template.Must(template.ParseFiles("templates/subtask_item.html"))
		tmpl.Execute(w, st)
	}
*/
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/kr_list.html",
		"templates/kr_item.html",
		//"templates/subtask_item.html",
		"templates/subtasks_view.html",
		"templates/subtask_edit_form.html",
	))
	krs, err := db.GetAllKeyResults()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.Execute(w, krs)
}
