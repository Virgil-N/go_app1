package login

import (
	"fmt"
	"html/template"
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
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
		/*
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
		*/
		//换成数据库，哈哈
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)
		c := session.DB("db1").C("user")
		var adminUser User
		c.Find(bson.M{"username": "tom"}).One(&adminUser)
		if adminUser.Username == username && adminUser.Password == password {
			//设置cookie
			expires := time.Now().AddDate(0, 0, 1)
			myCookie := http.Cookie{
				Name:    "testMyCookie",
				Value:   "success",
				Path:    "/",     //所有页面都可以读取cookie
				Expires: expires, //可加可不加？
				MaxAge:  86400,
			}
			http.SetCookie(w, &myCookie)
			fmt.Println(myCookie)
			fmt.Println(username, password)
			io.WriteString(w, "success")
		} else {
			io.WriteString(w, "fail")
		}

	}

}
