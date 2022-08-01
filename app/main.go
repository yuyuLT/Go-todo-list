package main

import (
  "database/sql"
  "fmt"
  "log"

  //インポート _にしないとコンパイルエラーになる。使用しない為
  _ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
    Name string
    Age  int
}


func main() {
      //1
    //DBを開く  なければ作成される
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    //終わったら閉じる
    defer DbConnection.Close()
    //DB作成 SQLコマンド
    cmd := `CREATE TABLE IF NOT EXISTS person(
        name STRING,
        age  INT)`

          //実行 結果は返ってこない為、_にする
    _, err := DbConnection.Exec(cmd)

     //エラーハンドリング
     if err != nil {
      fmt.Println("エラー")
      log.Fatalln(err)
  }

  cmd = "INSERT INTO person (name, age) VALUES (?, ?)"

  _, err = DbConnection.Exec(cmd, "Nancy", 25)
    if err != nil {
        log.Fatalln(err)
    }
    _, err = DbConnection.Exec(cmd, "Mike", 25)
    if err != nil {
        log.Fatalln(err)
    }

    //Update
    //データの更新 Mike が存在する場合
    cmd = "UPDATE person SET age = ? WHERE name = ?"
    _, err = DbConnection.Exec(cmd, 40, "Mike")
    if err != nil {
        log.Fatalln(err)
    }

    //Read
    //Get  All
    //マルチセレクト
    //データ全てをループで表示
    //Queryは全て取得する
    cmd = "SELECT * FROM person"
    rows, _ := DbConnection.Query(cmd)
    defer rows.Close()
    //structを作成
    var pp []Person
    //取得したデータをループでスライスに追加　for rows.Next()
    for rows.Next() {
        var p Person
        //scan データ追加
        err := rows.Scan(&p.Name, &p.Age)
        if err != nil {
            log.Println(err)
        }
        pp = append(pp, p)
    }
    err = rows.Err()
    if err != nil {
        log.Fatalln(err)
    }
    //表示
    for _, p := range pp {
        fmt.Println(p.Name, p.Age)
    }

    //特定のデータを取得
    cmd = "SELECT * FROM person where age = ?"
    //age = 20にして実行
    //QueryRowは最初の一件だけ取得する。
    row := DbConnection.QueryRow(cmd, 25)
    var p Person
    err = row.Scan(&p.Name, &p.Age)
    if err != nil {
        //データがなかったら
        if err == sql.ErrNoRows {
            log.Println("No row")
            //それ以外のエラー
        } else {
            log.Println(err)
        }
    }
    fmt.Println("セレクト→",p.Name, p.Age)
}