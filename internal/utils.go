package internal

import (
	"fmt"
	"os"
	"strconv"
)

// this makes me dislike golang...
func popIndex(values []string, delIndex int) []string {
	var new []string
	for index, value := range values {
		if index != delIndex {
			new = append(new, value)
		}
	}
	return new
}

// return true if it's a param, and the index of the param
func getParam(value string) (int, bool) {
	var isParam bool
	var index int64
	if len(value) > 1 && value[0] == '$' {
		isParam = true
		index, _ = strconv.ParseInt(value[1:], 10, 0)
	}
	return int(index), isParam
}

// return index if value exists, else -1
func find[T comparable](values []T, target T) int {
	for index, value := range values {
		if value == target {
			return index
		}
	}
	return -1
}

// format a list of items as a table and print it
func printTable(items []string) {
	var width int = 20
	for _, item := range items {
		if len(item) > width {
			item = item[:width-5] + "..."
		}
		paddedStr := fmt.Sprintf("%-*s", width, item)
		fmt.Printf("%s", paddedStr)
	}
	fmt.Printf("\n")
}

// returns default config path (~/.cly) or value in CLYPATH
func configPath() string {
	customPath := os.Getenv("CLYPATH")
	dirname, _ := os.UserHomeDir()
	if len(customPath) > 0 {
		return customPath
	}
	return dirname + "/.cly.yaml"
}
