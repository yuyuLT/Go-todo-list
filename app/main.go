package main

import (
	"example.com/module/data_controller"
	"fmt"
	"os"
)

func main() {

	fmt.Print("新規ユーザ登録:1 ログイン:2 > ")
	var status int
	fmt.Scan(&status)

	var id int

	if status == 1 {
		//新規登録
		fmt.Print("あなたのユーザ名を設定してください > ")
		var name string
		fmt.Scan(&name)

		fmt.Print("ログインID（整数）を設定してください > ")
		fmt.Scan(&id)

		fmt.Print("ログインパスワードを設定してください > ")
		var password string
		fmt.Scan(&password)

		//ユーザ情報を登録
		data_controller.RegisterUser(id, password, name)

	} else if status == 2 {
		//ログイン機能
		fmt.Print("ログインID（整数）を入力してください > ")
		fmt.Scan(&id)

		fmt.Print("ログインパスワードを入力してください > ")
		var password string
		fmt.Scan(&password)

		//ユーザ情報を取得
		data_controller.Login(id, password)
	} else {
		fmt.Println("値が不正です。半角数字1か2を入力してください")
		os.Exit(0)
	}

	for {
		//タスク一覧表示
		data_controller.ShowTasks(id)

		fmt.Print("タスク登録:1 タスク削除:2 ログアウト:3 > ")
		var action int
		fmt.Scan(&action)

		//タスク登録
		if action == 1 {
			fmt.Print("タスクを入力してください > ")
			var content string
			fmt.Scan(&content)

			data_controller.RegisterTask(id, content)
		}

		//タスク削除
		if action == 2 {
			fmt.Print("削除したいタスクの番号を入力してください > ")
			var task_num int
			fmt.Scan(&task_num)

			data_controller.DeleteTask(id, task_num)
		}

		//ログアウト処理
		if action == 3 {
			fmt.Println("ログアウトしました")
			os.Exit(0)
		}
	}
}
