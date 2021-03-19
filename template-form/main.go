package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // urlが渡すオプションを解析,POSTに対してはレスポンスパケットのボディを解析
	// ParseFormメソッドをコールしないと、以下でフォームのデータを取得できない
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!") // wに書き込まれたものがクライアントに出力される
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // リクエストを取得するメソッド
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// ログインデータがリクエストされ、データが表示される
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil) // 監視するポート設定
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
