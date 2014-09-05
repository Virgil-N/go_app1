package main

import (
	"controller/accountAction"
	"controller/login"
	"html/template"
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	t, _ := template.ParseFiles("./view/page/index.html")
	t.Execute(w, nil)
}

func main() {
	//静态文件处理函数，这些表明根目录是src？
	http.Handle("/css/", http.FileServer(http.Dir("view")))
	http.Handle("/js/", http.FileServer(http.Dir("view")))
	http.Handle("/img/", http.FileServer(http.Dir("view")))

	//动态文件处理函数
	http.HandleFunc("/", indexPage)
	//使用外部包函数
	http.HandleFunc("/login", login.Login)
	//记住自己加上Action的原因
	http.HandleFunc("/accountAction", accountAction.AccountAction)
	//注册服务并监听
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	log.Fatalln("listening")
}
