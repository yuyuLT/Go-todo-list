package data_controller

import (
	"fmt"
	"log"
	"database/sql"
	"example.com/module/db_operation"
)

type User struct{
	Id int
	Password string
	Name string
}

func CreateDb(){
	db_operation.CreateDb()
}

func InsertUser(id int, password string, name string){
	db_operation.InsertUser(id, password, name)
}

func SelectUser(id int){
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
    fmt.Println(user.Id, user.Password, user.Name)
}
