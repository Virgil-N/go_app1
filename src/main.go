package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Content-Type", "text/html")
	t, _ := template.ParseFiles("view/page/index.html")
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Add("Content-Type", "text/html")
		t, _ := template.ParseFiles("view/page/login.html")
		t.Execute(w, nil)

	} else {
		r.ParseForm()
		var temp = r.Form["username"][0] + r.Form["password"][0]
		w.Write([]byte(temp))
	}
}

func main() {

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	log.Fatalln("listening")
}
