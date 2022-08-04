module example.com/module

go 1.17

require github.com/mattn/go-sqlite3 v1.14.14

replace example.com/module/data_controller => ../data_controller

replace example.com/module/db_operation => ../db_operation
