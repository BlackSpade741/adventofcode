package main

import (
	"fmt"
	"strconv"
)

func day11() {
	input := ReadOneLineIntInput("day11input.txt", " ")
	
	fmt.Println(input)

	for i := 0; i < 75; i ++ {
		// fmt.Println("blink", i)

		newInput := make([]int, 0)

		for _, num := range(input) {
			if num == 0 {
				newInput = append(newInput, 1)
				// fmt.Println("replace 0 with 1")
			} else if len(strconv.Itoa(num)) % 2 == 0 {
				snum := strconv.Itoa(num) 
				halfl := len(snum) / 2
				snum1 := snum[0:halfl]
				snum2 := snum[halfl:]
				// fmt.Printf("split %s into %s and %s\n", snum, snum1, snum2)
				
				num1, _ := strconv.Atoi(snum1)
				num2, _ := strconv.Atoi(snum2)
				newInput = append(newInput, num1)
				newInput = append(newInput, num2)
			} else {
				newNum := num * 2024
				newInput = append(newInput, newNum)
				// fmt.Printf("Replace %d with %d\n", num, newNum)
			}
		}

		fmt.Println("num stones left:", len(newInput))
		// fmt.Println("ratio:", float64( len(newInput)) / float64(len(input)))
		input = newInput
		
	}
	// fmt.Println("num stones left:", len(input))
}