package main

import (
	"encoding/json"
	"net/http"
)

type fruit struct {
	Name  string `json:"hadfaname"`
	Price string `json:"ppppprice"`
}

func ReturnJson(w http.ResponseWriter, r *http.Request) {
	var s fruit
	s.Name = "apple"
	s.Price = "$2/kg"

	message, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application")
	w.Write(message)
}

func main() {
	http.HandleFunc("/login", ReturnJson)
	http.ListenAndServe(":9090", nil)
}
