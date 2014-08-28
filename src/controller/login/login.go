package login

import (
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
