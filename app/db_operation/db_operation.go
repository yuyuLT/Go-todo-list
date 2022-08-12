package db_operation

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func CreateTables(){
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS user(
		id INT PRIMARY KEY,
		password STRING,
		name STRING)`

	_, err := DbConnection.Exec(cmd)

	if err != nil {
        fmt.Println("エラー")
        log.Fatalln(err)
    }

	cmd = `CREATE TABLE IF NOT EXISTS task(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INT,
		date TIMESTAMP,
		content STRING)`

	_, err = DbConnection.Exec(cmd)

	if err != nil {
        fmt.Println("エラー")
        log.Fatalln(err)
    }
}

func InsertUser(id int, password string, name string){
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "INSERT INTO user (id, password, name) VALUES (?, ?, ?)"

	_, err := DbConnection.Exec(cmd, id, password, name)
    if err != nil {
        log.Fatalln(err)
		os.Exit(0)
    }
}

func SelectUser(id int) *sql.Row{
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "SELECT * FROM user WHERE id = ?"
	row := DbConnection.QueryRow(cmd, id)
    
	return row
}

func InsertPost(user_id int, content string){
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "INSERT INTO task (user_id, date, content) VALUES (?, CURRENT_TIMESTAMP, ?)"

	_, err := DbConnection.Exec(cmd, user_id, content)
    if err != nil {
        log.Fatalln(err)
		os.Exit(0)
    }
}

