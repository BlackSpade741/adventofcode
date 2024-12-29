package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func day9() {
	file, err := os.Open("day9input.txt")

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := scanner.Text()
	input := make([]int, 0)

	for _, c := range strings.Split(s, "") {
		value, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		input = append(input, value)
	}

	fmt.Println("input:", input)

	// loop through from the beginning, stop when we reach the same index
	// keep track of how many more blocks in the file we still need to move
	// keep track of the current index
	// for i := 0; i < backId * 2; i++ {
	// 	fmt.Println("i:", i, " size:", input[i])
		
	// 	if i % 2 == 0 {
	// 		fmt.Println("file, adding to checksum")
	// 		fileId := i / 2
			
	// 		for j := 0; j < input[i]; j ++ {
	// 			fmt.Println("Adding", blockPos, "*", fileId, "to checksum")
	// 			checksum += blockPos * fileId
	// 			blockPos += 1
	// 		}
	// 	} else {
	// 		fmt.Println("space")
			

			
	// 		for j := 0; j < input[i]; j ++ {
	// 			if numBlocksLeft <= 0 {
	// 				fmt.Println("move backId as weve run out of blocks to move")
	// 				backId -= 1
	// 				numBlocksLeft = input[backId * 2]
	// 				fmt.Println("backId:", backId, "numBlocksLeft:", numBlocksLeft)
	// 			}
	// 			fmt.Println("Adding", blockPos, "*", backId, "to checksum")
	// 			checksum += blockPos * backId
	// 			numBlocksLeft -= 1

	// 			blockPos += 1
	// 		}
	// 	}
	// }
	// for numBlocksLeft > 0 {
	// 	checksum += blockPos * backId
	// 	blockPos += 1
	// 	numBlocksLeft -= 1
	// }

	backId := len(input) / 2

	fmt.Println("Back id:", backId)

	numBlocksLeft := input[len(input) - 1]
	fmt.Println("num blocks left:", numBlocksLeft)
	// blockPos := 0

	checksum := 0

	// start from the very back and try to fit into a space
	// we can calculate that part of the checksum if we ended up moving it
	// start filling space from the beginning
	
	// use input to keep track of how many spaces are left 
	// consider going back to the beginning every time - also to calculate the block position


	fmt.Println("checksum:", checksum)
}