package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Form struct {
	Name string
	Age string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set( "Access-Control-Allow-Methods","POST" )

	if r.Method == "POST"{
		decoder := json.NewDecoder(r.Body)
		var f Form
		err:=decoder.Decode(&f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
		}
		js,err := json.Marshal(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
            return
		}
		w.Write(js)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
