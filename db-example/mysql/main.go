package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, err := sql.Open("mysql", "root:secret@/go-mysql?charset=utf8") // user:password@/(tcp)dbname?charset=utf8
	checkError(err)

	// 挿入
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkError(err)
	res, err := stmt.Exec("nozomu", "アルバイト", "2020-08-01")
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)

	// 更新
	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkError(err)

	res,err = stmt.Exec("nozomuupdate", id)
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
	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkError(err)

	res, err = stmt.Exec(id)
	checkError(err)

	affect, err = res.RowsAffected()
	checkError(err)

	fmt.Println(affect)

	db.Close()
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}