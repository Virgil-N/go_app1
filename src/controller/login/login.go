package login

import (
	// "fmt"
	"html/template"
	"io"
	// "labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"net/http"
)

type User struct {
	Username string
	Age      string
	Gender   string
	Password string
	Married  string
	Score    int
	Likes    []string
	Friends  []string
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "text/html")
		t, _ := template.ParseFiles("view/page/login.html")
		t.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		w.Header().Set("Content-Type", "text/plain")
		//暂时用“tom”，以后换数据库的数据
		if username == "tom" {
			io.WriteString(w, "matched")
		} else {
			io.WriteString(w, "not matched")
		}
	}

}
