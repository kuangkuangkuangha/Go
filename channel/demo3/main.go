package mian



func worker(sent chan int,get chan int){
	i:= <-sent
	get <- i*i
}


func main (){
	jobs := make(chan int,100)
	// results := make(chan int,100)

	for i:=0 ; i<5; i++{
		jobs <-i
	}
	close(jobs)
	
}