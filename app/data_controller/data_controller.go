package data_controller

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"example.com/module/db_operation"
)

type User struct{
	Id int
	Password string
	Name string
}

func CreateUserTable(){
	db_operation.CreateUserTable()
}

func InsertUser(id int, password string, name string){
	db_operation.InsertUser(id, password, name)
	fmt.Println("新規登録に成功しました")
}

func Login(id int, password string){
	_, pass, _ := SelectUser(id);
	if pass == password{
		fmt.Println("ログインに成功しました")
	}else{
		fmt.Println("ログインに失敗しました")
		os.Exit(0)
	}
};

func SelectUser(id int)(int,string,string){
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
