package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"kr-app/db"
	"kr-app/models"
)

// ========== Key Results ==========

func GetAllKRs(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, description FROM key_results")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var krs []models.KeyResult
	for rows.Next() {
		var kr models.KeyResult
		rows.Scan(&kr.ID, &kr.Title, &kr.Description)
		krs = append(krs, kr)
	}
	json.NewEncoder(w).Encode(krs)
}

func CreateKR(w http.ResponseWriter, r *http.Request) {
	var kr models.KeyResult
	json.NewDecoder(r.Body).Decode(&kr)
	res, err := db.DB.Exec("INSERT INTO key_results (title, description) VALUES (?, ?)", kr.Title, kr.Description)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	id, _ := res.LastInsertId()
	kr.ID = int(id)
	json.NewEncoder(w).Encode(kr)
}

func GetKR(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	row := db.DB.QueryRow("SELECT id, title, description FROM key_results WHERE id = ?", id)

	var kr models.KeyResult
	err := row.Scan(&kr.ID, &kr.Title, &kr.Description)
	if err != nil {
		http.Error(w, "KR no encontrado", 404)
		return
	}

	rows, _ := db.DB.Query("SELECT id, kr_id, title, done FROM sub_tasks WHERE kr_id = ?", id)
	defer rows.Close()
	for rows.Next() {
		var st models.SubTask
		rows.Scan(&st.ID, &st.KRID, &st.Title, &st.Done)
		kr.SubTasks = append(kr.SubTasks, st)
	}

	json.NewEncoder(w).Encode(kr)
}

func UpdateKR(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var kr models.KeyResult
	json.NewDecoder(r.Body).Decode(&kr)

	_, err := db.DB.Exec("UPDATE key_results SET title = ?, description = ? WHERE id = ?", kr.Title, kr.Description, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteKR(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.DB.Exec("DELETE FROM key_results WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ========== Sub-Tareas ==========

func GetSubTasks(w http.ResponseWriter, r *http.Request) {
	krID := mux.Vars(r)["kr_id"]
	rows, err := db.DB.Query("SELECT id, kr_id, title, done FROM sub_tasks WHERE kr_id = ?", krID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var subtasks []models.SubTask
	for rows.Next() {
		var st models.SubTask
		rows.Scan(&st.ID, &st.KRID, &st.Title, &st.Done)
		subtasks = append(subtasks, st)
	}
	json.NewEncoder(w).Encode(subtasks)
}

func CreateSubTask(w http.ResponseWriter, r *http.Request) {
	krIDStr := mux.Vars(r)["kr_id"]
	krID, _ := strconv.Atoi(krIDStr)

	var st models.SubTask
	json.NewDecoder(r.Body).Decode(&st)
	st.KRID = krID

	res, err := db.DB.Exec("INSERT INTO sub_tasks (kr_id, title, done) VALUES (?, ?, ?)", st.KRID, st.Title, st.Done)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	id, _ := res.LastInsertId()
	st.ID = int(id)
	json.NewEncoder(w).Encode(st)
}

func UpdateSubTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var st models.SubTask
	json.NewDecoder(r.Body).Decode(&st)

	_, err := db.DB.Exec("UPDATE sub_tasks SET title = ?, done = ? WHERE id = ?", st.Title, st.Done, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteSubTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.DB.Exec("DELETE FROM sub_tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
