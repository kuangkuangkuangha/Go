package main

import (
	"fmt"
	"time"
	
)

func hello(){
	fmt.Println("你好goroutine")
	for i := -1; i != 0; i-- {
		fmt.Println(i)
	}
}

func main(){
	go hello()

	for i := 1; i != 0; i++ {
		fmt.Println(i)
	}
	// fmt.Println("我在测试grouting")
	time.Sleep(time.Second)
}