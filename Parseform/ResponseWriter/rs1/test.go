package main

import (
	"fmt"
	"log"
	"net/http"
)

func noAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(401)
		w.Write([]byte("你是一个臭傻逼"))
		w.Write([]byte("------以下是你的请求信息-----"))
		fmt.Fprintln(w, "请求方法", r.Method)
		fmt.Fprintln(w, "请求主机", r.Host)
		fmt.Fprintln(w, "请求头", r.Header)
		fmt.Fprintf(w, "要经过认证才能访问这个接口")
	}
}

func main() {
	http.HandleFunc("/auth", noAuth)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println(err)
	}
}
