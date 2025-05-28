package models

type KeyResult struct {
	ID          int
	Title       string
	Description string
	SectorID    int
	SectorName  string
	SubTasks    []SubTask
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

type KrUpdPage struct {
	ID          int
	Title       string
	Description string
	SectorID    int
	Sectors     []Sector
}
