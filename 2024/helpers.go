package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func Print2DArray[T int|string](input [][]T) {
	for _, arr := range(input) {
		fmt.Println(arr)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadOneLineIntInput(inputFile string, sep string) []int {
	file, err := os.Open(inputFile)

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]int, 0)

	scanner.Scan()
	s := scanner.Text()

	for _, snum := range strings.Split(s, sep) {
		num, err := strconv.Atoi(snum)

		if err == nil {
			input = append(input, num)
		}
	}
	
	return input
}

func ReadMultiLineInput(inputFile string) []string {
	file, err := os.Open(inputFile)

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)

	for scanner.Scan() {
		s := scanner.Text()

		input = append(input, s)
	}
	
	return input
}

func Read2dArrayInput(inputFile string, sep string) [][]string {
	file, err := os.Open(inputFile)

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]string, 0)

	for scanner.Scan() {
		s := scanner.Text()
		line := make([]string, 0)

		for _, s := range strings.Split(s, sep) {
			line = append(line, s)
		}
		input = append(input, line)
	}
	
	return input
}