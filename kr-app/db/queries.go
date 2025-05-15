package db

import (
	//"database/sql"
	"kr-app/models"
	"log"
)

//var DB *sql.DB

// Key Results

func InsertKeyResult(kr models.KeyResult) (int, error) {
	res, err := DB.Exec("INSERT INTO key_results (title, description, sector_id) VALUES (?, ?, ?)", kr.Title, kr.Description, kr.SectorID)
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

func GetAllSectors() ([]models.Sector, error) {
	rows, err := DB.Query("SELECT id, name FROM sectors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sectors []models.Sector
	for rows.Next() {
		var sector models.Sector
		if err := rows.Scan(&sector.ID, &sector.Name); err != nil {
			return nil, err
		}
		sectors = append(sectors, sector)
	}
	return sectors, nil
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
	log.Printf("En DB.queries.UpdateSubTask, st: %v\n", st)
	_, err := DB.Exec("UPDATE sub_tasks SET title = ?, done = ? WHERE id = ?", st.Title, st.Done, st.ID)
	return err
}

func UpdateSubTaskTitle(id int, title string) error {
	query := `UPDATE subtasks SET title = ? WHERE id = ?`
	_, err := DB.Exec(query, title, id)
	return err
}

func DeleteSubTask(id int) error {
	_, err := DB.Exec("DELETE FROM sub_tasks WHERE id = ?", id)
	return err
}

func GetSubTasksByKRID(id int) ([]models.SubTask, error) {
	rows, err := DB.Query("SELECT * FROM sub_tasks WHERE kr_id = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subtks []models.SubTask

	for rows.Next() {
		var st models.SubTask

		err := rows.Scan(&st.ID, &st.KRID, &st.Title, &st.Done)

		if err != nil {
			return nil, err
		}

		log.Printf("SubTask: %v\n", st)

		subtks = append(subtks, st)
	}
	return subtks, nil
}
