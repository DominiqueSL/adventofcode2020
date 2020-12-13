package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func LineParser(line string) (int, int, string, string) {
	splitDash := strings.Split(line, "-")
	minOcc, err1 := strconv.Atoi(splitDash[0]) // first index of full length string contains minimum amount of occurrences
	if err1 != nil {
		panic(err1)
	}
	splitSpace := strings.Split(splitDash[1], " ") // split line by space
	maxOcc, err2 := strconv.Atoi(splitSpace[0]) // third index of full length string contains the maximum amount of occurrences string may contain
	if err2 != nil {
		panic(err2)
	}
	checkString := splitSpace[len(splitSpace) - 1] // "password" string which to check for certain charactes.
	charCheck := string(splitSpace[1][0])  // Get first index of second slice item to obtain the character for which the string has to be checked
	return minOcc, maxOcc, checkString, charCheck
}

func CheckPasswordPartTwo(inputFile string, delimiter string) int {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	splitString := strings.Split(string(data), delimiter)
	var firstIndex int
	var lastIndex int
	var checkString string
	var charCheck string

	counter := 0 // Initialize counter

	for _, line := range splitString {
		firstIndex, lastIndex, checkString, charCheck = LineParser(line)
		if charCheck == checkString[firstIndex - 1:firstIndex] && charCheck != checkString[lastIndex - 1: lastIndex]  {
			counter ++
		} else if charCheck != checkString[firstIndex - 1:firstIndex] && charCheck == checkString[lastIndex - 1: lastIndex]{
			counter ++
		}
	}
	return counter
}

func CheckPasswordPartOne(inputFile string, delimiter string) int {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	splitString := strings.Split(string(data), delimiter)
	var minOcc int
	var maxOcc int
	var checkString string
	var charCheck string

	counter := 0 // Initialize counter

	for _, line := range splitString {
		minOcc, maxOcc, checkString, charCheck = LineParser(line)
		counts := strings.Count(checkString, charCheck)
		if counts >= minOcc && counts <= maxOcc {
			counter ++
		}
	}
	return counter
}

func main()  {
	fmt.Println(CheckPasswordPartTwo("input", "\n"))
}