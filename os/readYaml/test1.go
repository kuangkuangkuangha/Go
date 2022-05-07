package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Folder map[string]string `yaml:"folder"`
}

func main() {
	var c conf

	con := c.getConf()
	fmt.Println(con)

	// 遍历map中的路径
	for _, key := range con.Folder {
		txt, _ := ioutil.ReadFile(key)

		content := string(txt)
		fmt.Println(content)
	}
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("/Users/zhangkuang/Documents/Desk/kuangkaungkaungha/Go-for-waiting/readYaml/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(ap1)
	return c
}
