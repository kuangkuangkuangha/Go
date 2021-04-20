package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func httpget(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//循环读取网页，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}

		result += string(buf[:n])
	}
	return
}

func working(start int, end int) {
	fmt.Printf("正在爬取第%d到%d页", start, end)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E8%B4%B4%E5%90%A7&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := httpget(url)
		if err != nil {
			fmt.Println("http get err", err)
			continue
		}
		fmt.Println("resulr=", result)
	}
}

func main() {
	//定义爬去起始页，终止页
	var start, end int
	fmt.Print("爬取的首页是")
	fmt.Scan(&start)

	fmt.Print("爬取的最后一页是")
	fmt.Scan(&end)

	working(start, end)
}
