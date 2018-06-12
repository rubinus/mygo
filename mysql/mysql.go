package mysql

import (
	"database/sql"

	"fmt"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/json-iterator/go"
)

func TestMysql() string {
	db, err := sql.Open("mysql", "root:root@/mysql?charset=utf8")
	checkErr(err)
	defer db.Close()

	//rows, _ := fetchRows(db, "SELECT * FROM t")
	//fmt.Println(*rows, "-------------")
	//for _, v := range *rows {
	//	fmt.Println(v["id"], v["c1"])
	//}

	/*aid, _ := insert(db, "INSERT INTO t( id,c1 ) VALUES( ?,? )", 3, "dd")
	fmt.Println(aid, "------aid----")
	row, _ := fetchRow(db, "SELECT * FROM t where id = ?", 3)
	fmt.Println(*row)*/

	//update
	//aid1, _ := exec(db, "UPDATE t SET c1 = ? WHERE id = ?", "hhhhh", 3)
	//fmt.Println(aid1, "------aid----")
	row1, _ := fetchRow(db, "SELECT * FROM user where User = ?", "root")
	fmt.Println(*row1)

	userJson, err := json.Marshal(*row1) //encoding/json
	fmt.Println(string(userJson))

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := jsonIterator.Marshal(*row1) //json_iterator

	return string(data)

}

type user struct {
	id int
	c1 string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//取一行数据，注意这类取出来的结果都是string
func fetchRow(db *sql.DB, sqlstr string, args ...interface{}) (*map[string]string, error) {
	stmtOut, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break //get the first row only
	}
	return &ret, nil
}

func fetchRows(db *sql.DB, sqlstr string, args ...interface{}) (*[]map[string]string, error) {
	stmtOut, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return &ret, nil
}

//插入
func insert(db *sql.DB, sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.LastInsertId()
}

//修改和删除
func exec(db *sql.DB, sqlstr string, args ...interface{}) (int64, error) {
	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	result, err := stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
	return result.RowsAffected()
}
