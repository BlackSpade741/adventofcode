package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func day5() {
	file, err := os.Open("day5input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	orderMap := make(map[string][]string)
	updates := make([][]string, 0)

	for scanner.Scan() {
		s := scanner.Text()
		// fmt.Println(s)

		if strings.Contains(s, "|") {
			result := strings.Split(s, "|")
			// fmt.Println(result)
			dependencies := orderMap[result[0]]

			orderMap[result[0]] = append(dependencies, result[1])
			// fmt.Println("orders")
		} else if strings.Contains(s, ",") {
			// fmt.Println("updates")
			result := strings.Split(s, ",")
			// fmt.Println(result)
			updates = append(updates, result)
		}
	}
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Println(orderMap)
	fmt.Println(updates)

	sum := 0

	for _, update := range updates {
		fmt.Printf("Processing update %s\n", update)
		seen := make(map[string]bool)
		updateLegal := true

		for _, item := range update {
			deps := orderMap[item]
			// fmt.Printf("Item %s, Dependencies: %s\n", item, deps)
			legal := true
			for _, dep := range deps {
				_, depseen := seen[dep]
				if depseen {
					// fmt.Printf("dependency %s has been seen before, illegal\n", dep)
					legal = false
				}
			}
			if legal {
				// fmt.Printf("item %s is safe\n", item)
			} else {
				// fmt.Printf("item %s is not safe\n", item)
				updateLegal = false
			}
			seen[item] = true
		}
		
		if updateLegal {
			fmt.Printf("update %s is legal, length: %d\n", update, len(update))
			// i := len(update) / 2
			// page, err := strconv.Atoi(update[i])
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(page)
			// sum = sum + page
		} else {
			fmt.Printf("update %s is illegal\n", update)

			// change order
			updateOrderMap := make(map[string][]string)

			for _, item := range update {
				for _, dep := range orderMap[item] {
					_, found := seen[dep]
					
					if found {
						updateOrderMap[item] = append(updateOrderMap[item], dep)
					}
				}
			}
			fmt.Println("order for this update:", updateOrderMap)

			
			// topological sort
			order := make([]string, 0)
			visited := make(map[string]bool)
						
			var dfs func(cur string)
			dfs = func(cur string) {
				// recurse then add to order
				_, hasVisited := visited[cur]

				if hasVisited {
					return
				}

				updateOrder := updateOrderMap[cur]
				for _, dep := range updateOrder {
					dfs(dep)
				}
				
				visited[cur] = true
				order = append(order, cur)
			}
			
			for _, item := range update {
				dfs(item)
			}

			fmt.Println("Updated order:", order)

			i := len(order) / 2
			page, err := strconv.Atoi(order[i])
			if err != nil {
				panic(err)
			}
			fmt.Println(page)
			sum = sum + page
		}
	}

	fmt.Println("Sum:", sum)
}


// func dfs()