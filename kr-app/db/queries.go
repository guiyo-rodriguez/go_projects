package db

import (
	//"database/sql"
	"kr-app/models"
)

//var DB *sql.DB

// Key Results

func InsertKeyResult(kr models.KeyResult) (int, error) {
	res, err := DB.Exec("INSERT INTO key_results (title, description) VALUES (?, ?)", kr.Title, kr.Description)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func GetAllKeyResults() ([]models.KeyResult, error) {
	rows, err := DB.Query("SELECT id, title, description FROM key_results")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var krs []models.KeyResult
	for rows.Next() {
		var kr models.KeyResult
		if err := rows.Scan(&kr.ID, &kr.Title, &kr.Description); err != nil {
			return nil, err
		}
		krs = append(krs, kr)
	}
	return krs, nil
}

func GetKeyResult(id int) (models.KeyResult, error) {
	var kr models.KeyResult
	err := DB.QueryRow("SELECT id, title, description FROM key_results WHERE id = ?", id).
		Scan(&kr.ID, &kr.Title, &kr.Description)
	return kr, err
}

func DeleteKeyResult(id int) error {
	_, err := DB.Exec("DELETE FROM key_results WHERE id = ?", id)
	return err
}

// Sub-Tareas

func InsertSubTask(st models.SubTask) (int, error) {
	res, err := DB.Exec("INSERT INTO sub_tasks (kr_id, title, done) VALUES (?, ?, ?)", st.KRID, st.Title, st.Done)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func GetSubTask(id int) (models.SubTask, error) {
	var st models.SubTask
	err := DB.QueryRow("SELECT id, kr_id, title, done FROM sub_tasks WHERE id = ?", id).
		Scan(&st.ID, &st.KRID, &st.Title, &st.Done)
	return st, err
}

func UpdateSubTask(st models.SubTask) error {
	_, err := DB.Exec("UPDATE sub_tasks SET title = ?, done = ? WHERE id = ?", st.Title, st.Done, st.ID)
	return err
}

func DeleteSubTask(id int) error {
	_, err := DB.Exec("DELETE FROM sub_tasks WHERE id = ?", id)
	return err
}
