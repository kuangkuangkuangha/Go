package main

import "fmt"

//开启两个goroutine，两个channel
//1.生成0～100的数字发送到 ch1
//2.从 ch1 中取出数据计算它的平方，把结果发送到 ch2 中

	func f1(ch chan int){
		for i :=0; i<100 ;i++{
			ch <- i
		}
		close(ch)
	}

	func f2(ch1 chan int, ch2 chan int){
		for {
			tmp, ok := <- ch1
			if !ok{
				break
			}
			ch2 <- tmp*tmp
		}
		close (ch2)
	}

//  chan<-   是一个只能发送的通道
//  <-chan  是一个只能接收的通道
func main(){
	ch1 := make (chan int , 100)
	ch2 := make (chan int , 200)

	go f1(ch1)
	go f2(ch1,ch2)

	for ret := range ch2{
		fmt.Println(ret)
	}
}