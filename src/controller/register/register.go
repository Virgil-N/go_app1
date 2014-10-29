package register

import (
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
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
	//password := r.FormValue("password")
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
	if adminUser.Username == username || username == "" {
		io.WriteString(w, "fail")
	} else if adminUser.Username != username && username != "" {
		io.WriteString(w, "success")
		//数据还没插入到数据库里面，上面只是演示一下
	}
}
