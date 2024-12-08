package main

import (
	"Advent-of-Code-2024/utils"
	"strings"
)

func DependenciesSort(array []int, assignments map[int][]int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if Contains(assignments[array[j]], array[j+1]) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Task1(assignments map[int][]int, updates [][]int) int {
	count := 0
	for _, update := range updates {
		isValidUpdate := true
		for i := 1; i < len(update); i++ {
			if !Contains(assignments[update[i-1]], update[i]) {
				isValidUpdate = false
			}
		}

		if isValidUpdate {
			count += update[len(update)/2]
		}
	}

	return count
}

func Task2(assignments map[int][]int, updates [][]int) int {
	count := 0
	for _, update := range updates {
		isValidUpdate := true
		for i := 1; i < len(update); i++ {
			if !Contains(assignments[update[i-1]], update[i]) {
				isValidUpdate = false
			}
		}

		if !isValidUpdate {
			update = DependenciesSort(update, assignments)
			count += update[len(update)/2]
		}
	}

	return count
}

func main() {
	lines := utils.ReadLines("input/day05.txt")
	var updates [][]int
	var assignments = make(map[int][]int)
	for _, line := range lines {
		if strings.Contains(line, ",") {
			fields := strings.Split(line, ",")

			var update []int
			for _, field := range fields {
				value, _ := utils.ToInt(field)
				update = append(update, value)
			}

			updates = append(updates, update)
		} else if strings.Contains(line, "|") {
			fields := strings.Split(line, "|")
			key, _ := utils.ToInt(fields[0])
			value, _ := utils.ToInt(fields[1])

			if assignments[key] != nil {
				assignments[key] = append(assignments[key], value)
			} else {
				assignments[key] = []int{value}
			}
		}
	}

	var r1 = Task1(assignments, updates)
	println(r1)
	var r2 = Task2(assignments, updates)
	println(r2)
}
