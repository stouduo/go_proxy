package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:9897"
	dbusername = "root"
	dbpassowrd = "root"
	dbname     = "test"
)

type mysql_db struct {
	db *sql.DB //定义结构体
}

func (f *mysql_db) mysql_open() { //打开
	Odb, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		fmt.Println("链接失败")
	}
	fmt.Println("链接数据库成功...........已经打开")
	f.db = Odb
}

func (f *mysql_db) mysql_close() { //关闭
	defer f.db.Close()
	fmt.Println("链接数据库成功...........已经关闭")
}

func (f *mysql_db) mysql_select(sql_data string) {
	rows, err := f.db.Query(sql_data)
	if err != nil {
		println(err)
	}
	for rows.Next() {
		var in_param string

		err = rows.Scan(&in_param)
		if err != nil {
			panic(err)
		}
		//user_id:=gojson.Json(in_param).Get("user_id").Tostring()
		fmt.Println(in_param)

	}
}

func main() {
	db := &mysql_db{}
	db.mysql_open()
	db.mysql_select("SELECT 1 FROM t   limit 10 ")
	db.mysql_close() //关闭
}
