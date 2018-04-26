package mysql

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysql() {
	db, err := sql.Open("mysql", "root:root@/phphbaseadmin?charset=utf8")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	if columns, err := rows.Columns(); err != nil {
		panic(err)
	} else {
		//拼接记录Map
		values := make([]sql.RawBytes, len(columns))
		scans := make([]interface{}, len(columns))

		for i := range values {
			scans[i] = &values[i]
		}
		//此处遍历在3W记录的时候，长达1分钟甚至更多
		for rows.Next() {
			_ = rows.Scan(scans...)
			each := map[string]interface{}{}
			for i, col := range values {
				each[columns[i]] = string(col)
			}
			fmt.Println(each)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
