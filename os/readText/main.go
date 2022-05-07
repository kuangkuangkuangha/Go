package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	txt, err := ioutil.ReadFile(`./新建文件.txt`)

	if err != nil {
		panic(err)
	}

	content := string(txt)

	fmt.Println(content)
}
