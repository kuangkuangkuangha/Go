package main

import (
	"fmt"
	"net/http"
	"text/template"
)
 

func handler(w http.ResponseWriter,r *http.Request){
	t,_:=template.ParseFiles("heool.html")
	t.Execute(w, "hello world")
}

func main(){
	http.HandleFunc("/hello",handler)
	http.ListenAndServe(":8080",nil)
	fmt.Println("nahao")
}