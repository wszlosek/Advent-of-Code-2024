package main

import (
	"Advent-of-Code-2024/utils"
	"math"
	"sort"
	"strings"
)

func Task1(list1 []int, list2 []int) int {
	var result = 0

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		result += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return result
}

func CountOccurrencesInt(list []int) map[int]int {
	counts := make(map[int]int)
	for _, item := range list {
		counts[item]++
	}
	return counts
}

func Task2(list1 []int, list2 []int) int {
	var result = 0
	var counts = CountOccurrencesInt(list2)

	for _, item := range list1 {
		result += item * counts[item]
	}

	return result
}

func main() {
	var list1 []int
	var list2 []int

	for _, line := range utils.ReadLines("input/day01.txt") {
		var parts = strings.Fields(line)
		var part1, _ = utils.ToInt(parts[0])
		var part2, _ = utils.ToInt(parts[1])

		list1 = append(list1, part1)
		list2 = append(list2, part2)
	}

	var r1 = Task1(list1, list2)
	println(r1)
	var r2 = Task2(list1, list2)
	println(r2)
}
