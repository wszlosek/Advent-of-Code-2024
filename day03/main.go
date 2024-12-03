package main

import (
	"Advent-of-Code-2024/utils"
	"regexp"
	"strings"
)

func Task1(lines []string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var matches [][][]string
	for _, line := range lines {
		matches = append(matches, re.FindAllStringSubmatch(line, -1))
	}

	return MulFromMatches(matches)
}

func Task2(lines []string) int {
	text := strings.Join(lines, "")
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	transformedText := re.ReplaceAllString(text, "")

	if strings.Contains(transformedText, "don't") {
		transformedText = strings.Split(transformedText, "don't")[0]
	}

	return Task1([]string{transformedText})
}

func MulFromMatches(matches [][][]string) int {
	result := 0
	for line := range matches {
		for group := range matches[line] {
			mul1, _ := utils.ToInt(matches[line][group][1])
			mul2, _ := utils.ToInt(matches[line][group][2])
			result += mul1 * mul2
		}
	}

	return result
}

func main() {
	lines := utils.ReadLines("input/day03.txt")

	var r1 = Task1(lines)
	println(r1)
	var r2 = Task2(lines)
	println(r2)
}
