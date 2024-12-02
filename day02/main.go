package main

import (
	"Advent-of-Code-2024/utils"
	"math"
	"strings"
)

func IsSafeReport(matrix []int) bool {
	isAscending := true
	isDescending := true

	for i := 0; i < len(matrix)-1; i++ {
		diff := math.Abs(float64(matrix[i]) - float64(matrix[i+1]))
		if diff < 1 || diff > 3 {
			return false
		} else if matrix[i] < matrix[i+1] {
			isDescending = false
		} else if matrix[i] > matrix[i+1] {
			isAscending = false
		}
	}

	return isAscending || isDescending
}

func Task1(matrix [][]int) int {
	result := 0
	for _, report := range matrix {
		if IsSafeReport(report) {
			result += 1
		}
	}

	return result
}

func Task2(matrix [][]int) int {
	result := 0
	for _, report := range matrix {
		if IsSafeReport(report) {
			result += 1
		} else {
			for i, _ := range report {
				newReport := utils.RemoveFromList(report, i)
				if IsSafeReport(newReport) {
					result += 1
					break
				}
			}
		}
	}

	return result
}

func main() {
	var matrix [][]int

	for i, line := range utils.ReadLines("input/day02.txt") {
		if len(matrix) <= i {
			matrix = append(matrix, []int{})
		}

		var parts = strings.Fields(line)
		for _, value := range parts {
			var part, _ = utils.ToInt(value)
			matrix[i] = append(matrix[i], part)
		}
	}

	var r1 = Task1(matrix)
	println(r1)
	var r2 = Task2(matrix)
	println(r2)
}
