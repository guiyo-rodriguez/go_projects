package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	//	_ "github.com/mattn/go-sqlite3" // o usa el driver de MySQL si prefieres
)

var DB *sql.DB

func Init() {
	var err error

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBCONN_STR")
	cfg.DBName = "recordings"

	fmt.Printf("cfg: %v\n", cfg)

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
		//log.Fatal(err)
		log.Printf("%v\n", err)
	}
	if _, err := DB.Exec(subTaskTable); err != nil {
		//log.Fatal(err)
		log.Printf("%v", err)
	}
}
