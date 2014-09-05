package accountAction

import (
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
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

func AccountAction(w http.ResponseWriter, r *http.Request) {
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
		//fmt.Println("search result: ", adminUser.Nickname)
		//现在就剩模板的部分了，哈哈
		//该死的模板(我错了，原来是路径搞的鬼！！！已经把account改成accountAction了，md查找原因花了1天时间！！)
		w.Header().Add("Content-Type", "text/html")
		t, _ := template.ParseFiles("view/page/account.html")
		//t.ExecuteTemplate(w, "tags", data)//为毛把原页面移除了？
		err = t.Execute(w, adminUser) //这个是ok的，嘿嘿
	}

}
