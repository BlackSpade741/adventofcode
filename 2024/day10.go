package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func day10() {
	file, err := os.Open("day10input.txt")

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([][]int, 0)

	starts := make([][2]int, 0)
	i := 0
	
	for scanner.Scan() {
		s := scanner.Text()
		line := make([]int, 0)

		for j, c := range strings.Split(s, "") {
			value, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			line = append(line, value)
			if value == 0 {
				starts = append(starts, [2]int{i, j})
			}
		}
		input = append(input, line)
		i += 1
	}

	Print2DArray(input)
	fmt.Println("starts:", starts)

	result := 0

	// dfs but only if the number is 1 larger
	// end at 9 and record
	// visited := make(map[[2]int]bool)
	// var visitedHeads map[[2]int] bool

	// var visited map[[2]int] bool


	var dfs func(curx int, cury int)
	dfs = func(curx int, cury int) {
		num := input[curx][cury]
		coord := [2]int{curx, cury}

		fmt.Printf("coord: %d num: %d\n", coord, num)

		if num == 9 {
			fmt.Println("Reached a trail end")
			
			// _, ok := visitedHeads[coord]

			// if ok {
			// 	fmt.Println("Already visited this one, dont count it")
			// } else {
			// 	fmt.Println("new trail end so count it")

			// 	result += 1
			// 	visitedHeads[coord] = true
			// }
			result += 1
		}

		// check left right up down
		deltas := [...][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		// visited[coord] = true
		
		for _, delta := range deltas {
			nextx := curx + delta[0]
			nexty := cury + delta[1]

			// _, hasVisited := visited[[2]int{nextx, nexty}]
			
			if nextx < 0 || nextx > len(input) - 1 || nexty < 0 || nexty > len(input[0]) - 1{
				continue
			}

			if input[nextx][nexty] == num + 1 {
				fmt.Printf("visit %d %d next\n", nextx, nexty)
				dfs(nextx, nexty)
			}
		}

		// delete(visited, coord)
	}
	
	for _, start := range(starts) {
		// visited = make(map[[2]int]bool)
		// visitedHeads = make(map[[2]int]bool)
		dfs(start[0], start[1])
	}

	fmt.Println("result:", result)
}