package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	// "strconv"
)

func day8() {
	file, err := os.Open("day8input.txt")

	if err != nil {
		fmt.Println("cannot read file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make([][]string, 0)
	antennae := make(map[string][][]int)
	lines := 0
	

	for scanner.Scan() {
		s := scanner.Text()
		input = append(input, strings.Split(s, ""))

		for i, c := range(strings.Split(s, "")) {
			if c != "." {
				antennae[c] = append(antennae[c], []int{lines, i})
			}
		}
		lines += 1
	}

	// fmt.Println(input)
	Print2DArray(input)
	fmt.Println(antennae)
	xmax := len(input) - 1
	ymax := len(input[0]) - 1

	// take each pair and create the places
	// antinodes := make(map[int]map[int]bool)

	for c, positions := range(antennae) {
		fmt.Println("character:", c)

		for i, pos1 := range(positions) {
			for j := i + 1; j < len(positions); j++ {
				pos2 := positions[j]

				fmt.Println("positions:", pos1, pos2)
				
				input[pos1[0]][pos1[1]] = "#"
				input[pos2[0]][pos2[1]] = "#"

				xdiff := pos1[0] - pos2[0]
				ydiff := pos1[1] - pos2[1]

				x1 := pos1[0] + xdiff
				y1 := pos1[1] + ydiff
				x2 := pos2[0] - xdiff
				y2 := pos2[1] - ydiff
				

				for 0 <= x1 && x1 <= xmax && 0 <= y1 && y1 <= ymax {
					fmt.Println(c,"antinode:", x1, y1)
					// _, ok1 := antinodes[x1]
					// if !ok1 {
					// 	antinodes[x1] = make(map[int]bool)
					// }
					// antinodes[x1][y1] = true
					input[x1][y1] = "#"
					
					x1 += xdiff
					y1 += ydiff
				}
				
				for 0 <= x2 && x2 <= xmax && 0 <= y2 && y2 <= ymax {
					fmt.Println(c,"antinode:", x2, y2)
					// _, ok2 := antinodes[x2]
					// if !ok2 {
					// 	antinodes[x2] = make(map[int]bool)
					// }
					// antinodes[x2][y2] = true
					input[x2][y2] = "#"

					x2 -= xdiff
					y2 -= ydiff
				}
			}
		}
	}
	// fmt.Println("antinodes:", antinodes)
	Print2DArray(input)
	total := 0

	for _, arr := range(input) {
		for _, c := range(arr) {
			if c == "#" {
				total += 1
			}
		}
	}

	fmt.Println("total:", total)
}
