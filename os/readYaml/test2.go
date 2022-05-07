package read

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type txt struct {
	txt map[string]string `yaml:"text_folder"`
}

func ReadText() map[string]string {
	var c txt
	txtMap := make(map[string]string, 10)

	con := c.getConfig()

	// 遍历map路径，取出内容
	for key, value := range con.txt {
		txt, _ := ioutil.ReadFile(value)
		content := string(txt)
		txtMap[key] = content
	}

	return txtMap
}

func (c *txt) getConfig() *txt {
	yamlFile, err := ioutil.ReadFile("../conf/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func main() {

}
