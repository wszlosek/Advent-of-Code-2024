package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	return lines
}

func ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func ToString(integer int) string {
	return strconv.Itoa(integer)
}

func GetKeys(m map[string]int) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func RemoveFromList(slice []int, i int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	return append(sliceCopy[:i], sliceCopy[i+1:]...)
}

func PrintList(list []int) {
	for _, item := range list {
		print(item, " ")
	}
	println()
}

func ReadMatrix(path string) [][]string {
	lines := ReadLines(path)
	var matrix [][]string

	for i, line := range lines {
		if len(matrix) <= i {
			matrix = append(matrix, []string{})
		}

		matrix[i] = strings.Split(line, "")
	}

	return matrix
}
