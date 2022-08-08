package main

import (
  "fmt"
  "example.com/module/data_controller"
)

func main() {

  //テーブルを作成
  data_controller.CreateUserTable()

  //新規登録
  fmt.Print("あなたのユーザ名を設定してください > ")
  var name string
  fmt.Scan(&name) 

  fmt.Print("ログインID（整数）を設定してください > ")
  var id int
  fmt.Scan(&id) 

  fmt.Print("ログインパスワードを設定してください > ")
  var password string
  fmt.Scan(&password) 

  //ユーザ情報を登録
  data_controller.InsertUser(id, password, name)
  
  //ユーザ情報を取得
  data_controller.SelectUser(id)
     
}