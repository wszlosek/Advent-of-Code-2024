package main

import (
	"Advent-of-Code-2024/utils"
)

func CheckDirection(matrix [][]string, x, y, dx, dy int, word string) bool {
	cols := len(matrix)
	rows := len(matrix[0])
	for i := 0; i < len(word); i++ {
		nx, ny := x+i*dx, y+i*dy
		if !IsValid(nx, ny, cols, rows) || matrix[nx][ny] != word[i:i+1] {
			return false
		}
	}
	return true
}

func IsValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func Task1(matrix [][]string) int {
	count := 0
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			for _, dir := range directions {
				if CheckDirection(matrix, r, c, dir[0], dir[1], "XMAS") {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	matrix := utils.ReadMatrix("input/day04.txt")
	var r1 = Task1(matrix)
	println(r1)
}
