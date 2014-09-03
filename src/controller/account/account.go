package account

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Age      int
	Gender   string
	Password string
	Article  []string
	Comment  []string
}

func Account(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	t, _ := template.ParseFiles("view/page/account.html")
	t.Execute(w, nil)

}
