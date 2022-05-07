package main

import "net/http"

func Redirt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	// w.Write([]byte("这是一个重定向请求")) // 不能使用w.write
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/login", Redirt)
	http.ListenAndServe(":9090", nil)
}
