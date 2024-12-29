package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func day7() {
	file, err := os.Open("day7input.txt")

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum_ := 0

	for scanner.Scan() {
		s := scanner.Text()

		parts := strings.Split(s, ": ")
		
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		var res []int

		values := strings.Split(parts[1], " ")
		
		for i, svalue := range(values) {
			value, err := strconv.Atoi(svalue)
			if err != nil {
				panic(err)
			}
			if i == 0 {
				res = []int{value}
				continue
			}
			// calculate by grabbing all results and adding the new part
			newres := make([]int, 0)
			for _, r := range(res) {
				newres = append(newres, r * value)
				newres = append(newres, r + value)

				concat, _ := strconv.Atoi(strconv.Itoa(r) + svalue)
				newres = append(newres, concat)
			}

			res = newres
		}

		for _, r := range(res) {
			if r == target {
				sum_ += target
				break
			}
		}
	}

	fmt.Println("sum:", sum_)
}