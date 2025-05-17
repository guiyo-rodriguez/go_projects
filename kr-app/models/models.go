package models

type KeyResult struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SectorID    int       `json:"sector_id"`
	SectorName  string    `json:"sector_name"`
	SubTasks    []SubTask `json:"sub_tasks,omitempty"`
}

type SubTask struct {
	ID    int    `json:"id"`
	KRID  int    `json:"kr_id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Sector struct {
	ID   int
	Name string
}

type PageDataKeyResults struct {
	KeyResults []KeyResult
	Sectors    []Sector
}

type KrEditPage struct {
	KeyResult KeyResult
	Sectors   []Sector
	Subtasks  []SubTask
}
