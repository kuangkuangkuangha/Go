package test

import (
	"encoding/json"
	"fmt"
)

func main() {
	type resp struct {
		id          int
		name        string
		cover       string
		view_count  string
		like_count  string
		duration    string
		author_name string
	}

	// b:=" {"id":289627667,"name":"拜登登机时不慎连摔三次，网友开启疯狂恶搞模式","cover":"http:/i0.hdslb.com/bfs/archive/78ce9ca6c83150eaa34b1bd109304bcf51260868.jpg","view_count":"455.7万","like_count":"26.7万","duration":"1:33","author_name":"观察者网","author_id":10330740,"bvid":"BV1Nf4y1t7G9","sort":"hot","filt":0} "
	var p2 resp
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
