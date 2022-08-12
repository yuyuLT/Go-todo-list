package main

import (
  "fmt"
  "example.com/module/data_controller"
)

func main() {

  fmt.Print("新規ユーザ登録:1 ログイン:2 > ")
  var status int
  var id int
  fmt.Scan(&status)
  
  if status == 1 {
    //新規登録
    data_controller.CreateTables()

    fmt.Print("あなたのユーザ名を設定してください > ")
    var name string
    fmt.Scan(&name) 

    fmt.Print("ログインID（整数）を設定してください > ")
    fmt.Scan(&id) 

    fmt.Print("ログインパスワードを設定してください > ")
    var password string
    fmt.Scan(&password) 

    //ユーザ情報を登録
    data_controller.InsertUser(id, password, name)

  }else if status == 2 {
    //ログイン機能
    fmt.Print("ログインID（整数）を入力してください > ")
    fmt.Scan(&id) 

    fmt.Print("ログインパスワードを入力してください > ")
    var password string
    fmt.Scan(&password) 

    //ユーザ情報を取得
    data_controller.Login(id,password)
  }else{
    fmt.Println("値が不正です。半角数字1か2を入力してください")
  }

  //タスク一覧表示
  data_controller.ShowPost(id)  

  //新規投稿登録
  fmt.Print("タスクを入力してください > ")
  var content string
  fmt.Scan(&content) 

  data_controller.InsertPost(id,content)  
}