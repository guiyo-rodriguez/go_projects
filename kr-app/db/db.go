package db

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
//	"os"

//	_ "github.com/mattn/go-sqlite3" // o usa el driver de MySQL si prefieres
)

var DB *sql.DB

func Init() {
	var err error
	
	cfg := mysql.NewConfig()
    //cfg.User = os.Getenv("DBUSER")
	cfg.User = "guille"
    //cfg.Passwd = os.Getenv("DBPASS")
	cfg.Passwd = "6z6b6ch3"
    cfg.Net = "tcp"
    cfg.Addr = "192.168.0.5:3306"
    cfg.DBName = "recordings"
	
	DB, err = sql.Open("sqlite3", "krapp.db") // usa "user:pass@tcp(localhost:3306)/dbname" para MySQL
	DB, err = sql.Open("mysql", cfg.FormatDSN()) 
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	krTable := `
	CREATE TABLE key_results (
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(80) NOT NULL,
		description VARCHAR(250)
	);`

	subTaskTable := `
	CREATE TABLE IF NOT EXISTS sub_tasks (
		id INT PRIMARY KEY AUTO_INCREMENT,
		kr_id INT NOT NULL,
		title VARCHAR(80) NOT NULL,
		done BOOLEAN DEFAULT FALSE,
		FOREIGN KEY (kr_id) REFERENCES key_results(id) ON DELETE CASCADE
	);`

	if _, err := DB.Exec(krTable); err != nil {
		log.Fatal(err)
	}
	if _, err := DB.Exec(subTaskTable); err != nil {
		log.Fatal(err)
	}
}
