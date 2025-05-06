package models

type KeyResult struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	SubTasks    []SubTask  `json:"sub_tasks,omitempty"`
}

type SubTask struct {
	ID     int    `json:"id"`
	KRID   int    `json:"kr_id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
}
