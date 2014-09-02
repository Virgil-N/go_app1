package login

import (
	"fmt"
	"html/template"
	"io"
	// "labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	"net/http"
	"time"
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
		password := r.FormValue("password")
		w.Header().Set("Content-Type", "text/plain")
		//暂时用“tom”，以后换数据库的数据
		if username == "tom" && password == "12345" {
			//添加cookie(注意添加cookie必须放在服务器返回数据之前)
			expires := time.Now().AddDate(0, 0, 1)
			myCookie := http.Cookie{
				Name:    "testMyCookie",
				Value:   "hello everyone",
				Path:    "/login",
				Expires: expires, //可加可不加？
				MaxAge:  86400,
			}
			http.SetCookie(w, &myCookie)
			fmt.Println(myCookie)

			io.WriteString(w, "matched")
		} else {
			io.WriteString(w, "not matched")
		}
	}

}
