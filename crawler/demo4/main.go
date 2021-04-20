package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/tidwall/gjson"
)


type resp struct {
	
	Channel_id  int 
	View_count  string
	Like_count  string
	Author_name string
	Author_id   string
	Bvid        string
	Card_type   string
}




func httpget(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	// defer resp.Body.Close()

	//循环读取网页，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			// fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		} //

		result += string(buf[:n])
	}
	return
}


// func spider(){

// 	url2 := strconv.Itoa(k)
// 	url := "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id="+url2+"&sort_type=hot&page_size=30"
// 	result, err := httpget(url)
// 	if err != nil {
// 		fmt.Println("http get err", err)

// 	}



// }






func main() {

	// temp := make([]int, 500)
	temp := make( chan int, 100)
	for k:=0 ;k<=100;k++{

	num := strconv.Itoa(k)

	url := "https://api.bilibili.com/x/web-interface/web/channel/category/channel_arc/list?id="+num

	result, err := httpget(url)
	if err != nil {
		fmt.Println("http get err", err)
	}
	

	for i := 0; i <= 5; i++ {

		num2 := strconv.Itoa(i)

		s_id := gjson.Get(result, "data.archive_channels."+num2+".id").String()

		id, _ := strconv.Atoi(s_id)

		// temp = append(temp, id)
		temp <- id
		
	}
	

}

for i:=0;i<=5;i++{
haha :=  <-temp
url2 := strconv.Itoa(haha)
url := "https://api.bilibili.com/x/web-interface/web/channel/multiple/list?channel_id="+url2+"&sort_type=hot&page_size=30"


result2, err := httpget(url)
	if err != nil {
		fmt.Println("http get err", err)
	}
	
	te := resp{
		
		Channel_id:  1,
		View_count:  gjson.Get(result2, "data.list.2.view_count").String(),
		Like_count:  gjson.Get(result2, "data.list.2.like_count").String(),
		Author_name: gjson.Get(result2, "data.list.2.author_name").String(),
		Author_id:   gjson.Get(result2, "data.list.2.author_id").String(),
		Bvid:        gjson.Get(result2, "data.list.2.bvid").String(),
		Card_type:   gjson.Get(result2, "data.list.2.card_type").String(),
	}

	fmt.Println(te)
	fmt.Println(temp)
}
}