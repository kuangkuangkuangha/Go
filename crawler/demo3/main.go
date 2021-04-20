package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

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

func getchannel() {

	for k := 0; k <= 100; k++ {
		num := strconv.Itoa(k)
		url := "https://api.bilibili.com/x/web-interface/web/channel/category/channel_arc/list?id=" + num
	    
		result, err := httpget(url)
	    if err != nil {
		fmt.Println("http get err", err)

		temp := make([]int, 100)

		for i := 0; i <= 5; i++ {

			num2 := strconv.Itoa(i)

			s_id := gjson.Get(result, "data.archive_channels."+num2+".id").String()

			id , _:= strconv.Atoi(s_id)
			
			temp := append(temp, id)

			fmt.Println(temp)

		}

	}

	
	}
	

}
func main(){

	getchannel()


}