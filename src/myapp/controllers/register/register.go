package register

import (
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

//下面的结构体声明每个controller都要加？
type User struct {
	Username string
	Age      int
	Gender   string
	Nickname string
	Password string
	Article  []string
	Comment  []string
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	w.Header().Set("Content-Type", "text/plain")

	//使用数据库
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("db1").C("user")
	var adminUser User
	c.Find(bson.M{"username": username}).One(&adminUser)
	//下面这样会出错，是个golang的bug
	// if adminUser.Username == username {
	// 	return "fail"
	// } else if adminUser.username != username && username != "" {
	// 	return "succes"
	// }
	if username == "" || password == "" {
		io.WriteString(w, "noValue")
	} else if adminUser.Username == username {
		io.WriteString(w, "fail")
	} else if adminUser.Username != username && username != "" {
		//怎么做啊？？？
		//c.Insert({"username": username, "nickname": "op", "password": "12345", "age": 19, "gender": "M", "article": ["hah", "flow me"], "comment": ["yes", "ok", "not bad"]})
		err = c.Insert(&User{username, 19, "M", "nick", "12345", []string{"hah", "flow me"}, []string{"yes", "ok", "not bad"}})
		if err != nil {
			io.WriteString(w, "error")
		}
		//注册成功添加cookie
		expires := time.Now().AddDate(0, 0, 1)
		myCookie := http.Cookie{
			Name:    "testMyCookie",
			Value:   username,
			Path:    "/",     //所有页面都可以读取cookie
			Expires: expires, //可加可不加？
			MaxAge:  86400,
		}
		http.SetCookie(w, &myCookie)
		io.WriteString(w, "success")

	}
}
