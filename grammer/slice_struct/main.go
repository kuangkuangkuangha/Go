package main

import "fmt"


type student struct{

	id int 
	name string
	password string

	}

func main(){
	
	fmt.Println("naihao")
		
		temp := make([]*student, 100)

		temp[0]=&student{
			id : 11,
			name :"xiaoming",
			password :"12345",
		}

		temp[1] = &student{
			id :22,
			name :"xiaohua",
			password:"gun",
		}

		for i:=2; i<=99; i++{
			temp[i]=&student{
				id : i,
				name: "zhangkuang" ,
				password:"gogogogo",
			}

		}

		fmt.Println(temp[0],temp[1])
		

		// for key,valve := range temp{
		// 	fmt.Println(key,valve)

		// }

		// fmt.Println("这是不带key值的")
		// for _,value:= range temp{
		// 	fmt.Println(value)
		// }





	}