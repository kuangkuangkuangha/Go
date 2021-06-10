package model

import (
	
	"fmt"
	"strconv"
)

func GetName(temp string )(id []int){

	// var id int
	// var player Player
	// var player2 []
	var shu int
	inter, _:= strconv.Atoi(temp)

	for inter > 0{
		shu = inter %10
		inter = inter/10
		
		 
		// DB.Table("players").Where("id = ?",id).Find(&player)
		id = append(id,shu)
		fmt.Println(id)
	}

	fmt.Println(id)


	// for _, value := range player2 {
	// 	name = append(name, value.Name)
	// }

	return id

}