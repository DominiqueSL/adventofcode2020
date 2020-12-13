package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func FindSum(input []int) int {
	// It should always return two entries.
	for _, numberZeroOppa := range input {
		for _, numberOneOppa := range input {
			for _, numberTwoOppa := range input {
				if numberZeroOppa + numberOneOppa+numberTwoOppa == 2020 {
					return numberZeroOppa * numberOneOppa * numberTwoOppa
				}
			}
		}
	}
	return 0
}

func FileParser(inputFile string) []int {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	splitString := strings.Split(string(data), "\n")
	var outArray = []int{}
	for _, number := range splitString {
		j, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		outArray = append(outArray, j)
	}
	return outArray
}

func main() {
	// Input should be a list of numbers.
	input_array := FileParser("input")
	//input_array := []int{556, 2984, 1721, 979, 366, 299, 675, 1456}
	println(FindSum(input_array))
}
