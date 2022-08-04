package main

import (
  "example.com/module/data_controller"
)

func main() {
  
  //テーブルを作成
  data_controller.CreateDb()

  //仮設定
  id := 101
  password := "password1"
  name := "名前"

  //ユーザ情報を登録
  data_controller.InsertUser(id, password, name)

  //ユーザ情報を取得
  data_controller.SelectUser(id)
     
}