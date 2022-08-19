package db_operation

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var DbConnection *sql.DB

func CreateTables() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS user(
		id INTEGER PRIMARY KEY,
		password TEXT,
		name TEXT)`

	_, err := DbConnection.Exec(cmd)

	if err != nil {
		fmt.Println("エラー")
		log.Fatalln(err)
	}

	cmd = `CREATE TABLE IF NOT EXISTS task(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		date TIMESTAMP,
		content TEXT)`

	_, err = DbConnection.Exec(cmd)

	if err != nil {
		fmt.Println("エラー")
		log.Fatalln(err)
	}
}

func InsertUser(id int, password string, name string) {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "INSERT INTO user (id, password, name) VALUES (?, ?, ?)"

	_, err := DbConnection.Exec(cmd, id, password, name)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
}

func SelectUser(id int) *sql.Row {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "SELECT * FROM user WHERE id = ?"
	row := DbConnection.QueryRow(cmd, id)

	return row
}

func GetTasks(id int) *sql.Rows {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "SELECT * FROM task WHERE user_id = ?"
	rows, _ := DbConnection.Query(cmd, id)

	return rows
}

func InsertTask(user_id int, content string) {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "INSERT INTO task (user_id, date, content) VALUES (?, CURRENT_TIMESTAMP, ?)"

	_, err := DbConnection.Exec(cmd, user_id, content)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
}

func DeleteTask(id int, task_num int) {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "DELETE FROM task WHERE id = (SELECT id FROM task WHERE user_id = ? LIMIT ?,1)"
	_, err := DbConnection.Exec(cmd, id, task_num)
	if err != nil {
		log.Fatalln(err)
	}
}
