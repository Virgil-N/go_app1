package account

import (
	"fmt"
	"html/template"
	// "io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"os"
)

type User struct {
	Username string
	Age      int
	Gender   string
	Nickname string
	Password string
	Article  []string
	Comment  []string
}

func Account(w http.ResponseWriter, r *http.Request) {
	myCookie, _ := r.Cookie("testMyCookie")
	myCookieValue := myCookie.Value
	if len(myCookieValue) > 0 {
		fmt.Println("get cookie success: ", myCookieValue)
		//连接数据库，获取用户的信息
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, false)
		c := session.DB("db1").C("user")
		var adminUser User
		c.Find(bson.M{"username": myCookieValue}).One(&adminUser)
		fmt.Println("search result: ", adminUser.Nickname)
		//现在就剩模板的部分了，哈哈
		//该死的模板
		w.Header().Add("Content-Type", "text/html")
		data := User{
			Username: "alice",
			Age:      24,
			Gender:   "F",
			Nickname: "aa",
			Password: "12345",
			Article:  []string{"hi", "o"},
			Comment:  []string{"hello", "tips"},
		}
		tpl := template.New("")
		tpl.Parses("view/page/account.html")
		tpl.Execute(w, data)
	}
	// w.Header().Add("Content-Type", "text/html")
	// t, err := template.ParseFiles("view/page/account.html")
	// t.Execute(w, nil)

}
