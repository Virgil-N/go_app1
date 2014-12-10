package main

import (
	"html/template"
	"log"
	"myapp/controllers/accountAction"
	"myapp/controllers/login"
	"myapp/controllers/pagination"
	"myapp/controllers/register"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	t, _ := template.ParseFiles("view/template/index.html")
	t.Execute(w, nil)
}

func main() {
	//静态文件处理函数
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/img/", http.FileServer(http.Dir("public")))

	//动态文件处理函数
	http.HandleFunc("/", indexPage)
	//使用外部包函数
	http.HandleFunc("/login", login.Login)
	//记住自己加上Action的原因
	http.HandleFunc("/accountAction", accountAction.AccountAction)
	//注册
	http.HandleFunc("/register", register.Register)
	//分页
	http.HandleFunc("/prevPage", pagination.PrevPage)
	//为毛会报错？(因为没清除pkg目录下的文件)
	http.HandleFunc("/nextPage", pagination.NextPage)
	//跳转页面
	http.HandleFunc("/goToPage", pagination.GoToPage)
	//注册服务并监听
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	log.Fatalln("listening")
}
