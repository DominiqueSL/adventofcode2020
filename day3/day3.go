package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func TraversePath(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	treeLines := strings.Split(string(data), "\n")
	return treeLines
}

func main()  {
	fmt.Println(TraversePath("input"))
}