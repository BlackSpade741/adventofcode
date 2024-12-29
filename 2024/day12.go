package main

import (
	"fmt"
	// "strconv"
)

func day12() {
	input := Read2dArrayInput("day12input.txt", "")
	
	Print2DArray(input)

	visited := make(map[[2]int]bool)

	var dfs func(i int, j int) [2]int

	dfs = func(i int, j int) [2]int {
		plant := input[i][j]
		deltas := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		loc := [2]int{i, j}
		visited[loc] = true
		area := 1
		perimeter := CalcFence(input, i, j)

		for _, delta := range deltas {
			x := i + delta[0]
			y := j + delta[1]
			nextloc := [2]int{x, y}

			_, hasVisited := visited[nextloc]

			if 0 > x || x >= len(input) || 0 > y || y >= len(input[0]) {
				continue
			}

			if input[x][y] == plant && !hasVisited {
				res := dfs(x, y)
				area += res[0]
				perimeter += res[1]
			}
		}
		return [2]int{area, perimeter}
	}

	price := 0
	for i, row := range input {
		for j, plant := range row {
			loc := [2]int{i, j}
			
			_, hasVisited := visited[loc]
			
			if !hasVisited {
				fmt.Println("visiting", loc, plant)
				res := dfs(i, j)
				fmt.Println("area:", res[0], "perimeter:", res[1])
				price += res[0] * res[1]
			}
			// fmt.Println(CalcFence(input, i, j))
		}
	}
	fmt.Println("price:", price)
}

func CalcFence(farm [][]string, i int, j int) int {
	deltas := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	plant := farm[i][j]
	total := 0
	
	for _, delta := range deltas {
		x := i + delta[0]
		y := j + delta[1]
		// fmt.Println(x, y)
		
		if 0 > x || x >= len(farm) || 0 > y || y >= len(farm[0]) {
			total += 1
			// fmt.Println("adding a fence as we're at the edge of the farm")
		} else if farm[x][y] != plant {
			total += 1
			// fmt.Println("adding a fence as the adjacent plant is different")
		}
	}
	return total
}
