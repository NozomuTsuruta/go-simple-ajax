package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=root password=secret dbname=go-postgres sslmode=disable")
	checkError(err)

	// 挿入
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkError(err)
	res, err := stmt.Exec("nozomu", "アルバイト", "2020-08-01")
	checkError(err)

	// PostgresはMySQLのインクリメンタルなIDのようなものがないため、この関数をサポートしていない
	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)

	// 更新
	stmt, err = db.Prepare("UPDATE userinfo SET username=$1 WHERE uid=$2")
	checkError(err)

	res, err = stmt.Exec("nozomuupdate", 1)
	checkError(err)

	affect, err := res.RowsAffected()
	checkError(err)

	fmt.Println(affect)

	// 検索
	rows, err := db.Query("SELECT * FROM userinfo")
	checkError(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkError(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 削除
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=$1")
	checkError(err)

	res, err = stmt.Exec(1)
	checkError(err)

	affect, err = res.RowsAffected()
	checkError(err)

	fmt.Println(affect)

	db.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
