package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	// "strconv"
)

func day6() {
	file, err := os.Open("day6input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maze := make([][]string, 0)
	var curx int
	var cury int

	for scanner.Scan() {
		s := scanner.Text()
		maze = append(maze, strings.Split(s, ""))

		for j, symbol := range s {
			if symbol == '^' {
				fmt.Println("Start has been found")
				curx = len(maze) - 1
				cury = j
			}
		}
	}
	fmt.Println(maze)
	fmt.Println("Start:", curx, cury)

	steps := 1
	xdir := -1
	ydir := 0
	options := 0
	xrocks := make(map[int][][]int)
	yrocks := make(map[int][][]int)

	// go through and walk
	for {
		// fmt.Println(curx, cury)
		// try to make step in the direction and check whether we can go
		nextx := curx + xdir
		nexty := cury + ydir

		if nextx < 0 || nextx >= len(maze) || nexty < 0 || nexty >= len(maze[0]) {
			fmt.Println("Exiting maze at", nextx, nexty)
			break
		}

		if maze[nextx][nexty] == "#" {
			fmt.Println("Rock at", nextx, nexty, "turn 90 degrees")

			xrocks[nextx] = append(xrocks[nextx], []int{nextx, nexty})
			yrocks[nexty] = append(yrocks[nexty], []int{nextx, nexty})

			fmt.Println("xrocks:", xrocks)
			fmt.Println("yrocks:", yrocks)

			xdir, ydir = getNextDir(xdir, ydir)

			fmt.Println("Current dir:", xdir, ydir)
			continue
		}

		// take a step and set maze if applicable
		curx = nextx
		cury = nexty
		if maze[curx][cury] == "." {
			maze[curx][cury] = "^"
			steps += 1


			// check if there is any obstacle we've already encoutnered in the next direction
			xrockshere := xrocks[curx]
			if len(xrockshere) > 0 {
				fmt.Println("xrocks:", xrockshere)
			}
			yrockshere := yrocks[cury]
			if len(yrockshere) > 0 {
				fmt.Println("yrocks:", yrockshere)
			}
			
		}
	}
	
	fmt.Println("steps:", steps)
	fmt.Println("options:", options)
}

func getNextDir(xdir int, ydir int) (int, int) {
	if xdir == -1 && ydir == 0 {
		// up to right
		return 0, 1
	}
	if xdir == 0 && ydir == 1 {
		// right to down
		return 1, 0
	}
	if xdir == 1 && ydir == 0 {
		// down to left
		return 0, -1
	}
	if xdir == 0 && ydir == -1 {
		return -1, 0
	}
	return 0, 0
}