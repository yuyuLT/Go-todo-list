package data_controller

import (
	"database/sql"
	"example.com/module/db_operation"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Id       int
	Password string
	Name     string
}

type Task struct {
	Id      int
	UserId  int
	Date    time.Time
	Content string
}

func CreateTables() {
	db_operation.CreateTables()
}

func InsertUser(id int, password string, name string) {
	db_operation.InsertUser(id, password, name)
	fmt.Println("新規登録に成功しました")
}

func Login(id int, password string) {
	_, pass, _ := SelectUser(id)
	if pass == password {
		fmt.Println("ログインに成功しました")
	} else {
		fmt.Println("ログインに失敗しました")
		os.Exit(0)
	}
}

func SelectUser(id int) (int, string, string) {
	row := db_operation.SelectUser(id)
	var user User
	err := row.Scan(&user.Id, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	return user.Id, user.Password, user.Name
}

func ShowPost(id int) {
	rows := db_operation.GetTasks(id)
	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.UserId, &task.Date, &task.Content)
		if err != nil {
			log.Println(err)
		}
		tasks = append(tasks, task)
	}
	err := rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	for i, task := range tasks {
		number := i + 1
		fmt.Println(number, task.Content)
	}
}

func InsertPost(id int, content string) {
	db_operation.InsertPost(id, content)
	fmt.Println("タスクを登録しました")
}

func DeletePost(id int, task_num int) {
	db_operation.DeletePost(id, task_num-1)
	fmt.Println("タスクを削除しました")
}
