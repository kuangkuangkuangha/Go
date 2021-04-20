package main

import "fmt"

func main(){
	var ch1 chan int 
	ch1 = make(chan int ,2)//带缓冲区通道，又称为同步通道
	ch1 = make(chan int)//无缓冲区通道，又称为异步通道
	ch1 <- 10
	x := <- ch1
	fmt.Println(x)
	close(ch1)
}