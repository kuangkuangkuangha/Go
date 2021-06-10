package model

import (
	
	"fmt"
	"log"
)

func GetRole(username string) int {
	var temp User
	if err := DB.Where("username = ?", username).Find(&temp).Error; err != nil {
		log.Println(err)
	}

	fmt.Println(temp.ROle)
	return temp.ROle

}
