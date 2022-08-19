
package db_operation

import (
	"database/sql"
    "gorm.io/driver/sqlite"
  	"gorm.io/gorm"
)

type User struct {
    Id       int `gorm:"primary_key"`
    Password string
    Name     string
}

type Task struct {
    gorm.Model
    UserId  int
    Content string
}

func CreateTables() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Task{})
}

func InsertUser(id int, password string, name string) {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

    var user = User{Id: id, Password: password, Name: name}
    db.Create(&user)
}

func SelectUser(id int) *sql.Row {
	var user User
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	row := db.Where("id = ?", id).First(&user).Row()
	return row
}

func GetTasks(id int) *sql.Rows {
	var task Task
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	rows, _ := db.Select("content").Where("user_id = ?", id).Find(&task).Rows()
	return rows
}

func InsertTask(user_id int, content string) {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	var task = Task{UserId: user_id, Content: content}
	db.Create(&task)
}

func DeleteTask(id int, task_num int) {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	db.Exec("DELETE FROM tasks WHERE id = (SELECT id FROM tasks WHERE user_id = ? LIMIT ? , 1)",id,task_num)
}
