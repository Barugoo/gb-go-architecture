package main

import (
	"fmt"
	"homework/binary"
	"strings"
)

func main() {
	intList := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	stringList := []string{"арбуз", "банан", "вишня", "груша", "дыня", "ежевика", "яблоко"}

	var searchValue = 10

	fmt.Printf("Found in index %d\n", binary.Search(0, len(intList), func(index int) int {
		if intList[index] == searchValue {
			return 0
		} else if intList[index] < searchValue {
			return -1
		} else {
			return 1
		}
	}))
	var searchStringValue = "дыня"

	fmt.Printf("Found in index %d\n", binary.Search(0, len(stringList), func(index int) int {
		return strings.Compare(stringList[index], searchStringValue)
	}))
}
