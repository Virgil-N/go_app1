package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var responseText = "ok"
		w.Write([]byte(responseText))
	})

	http.ListenAndServe(":8080", nil)
	log.Fatalln("listening")
}
