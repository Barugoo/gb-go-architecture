package main

import (
	"fmt"
)

func Search(a []int, search int) (result int, count int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > search:
		result, count = Search(a[:mid], search)
	case a[mid] < search:
		result, count = Search(a[mid+1:], search)
		result += mid + 1
	default:
		result = mid
	}
	count++
	return
}

func main() {

	searchArray := []int{1, 4, 8, 13, 18, 21, 29, 69, 77, 99}

	fmt.Printf("Масив: %v\n", searchArray)
	for _, num := range searchArray {
		result, count := Search(searchArray, num)
		fmt.Printf("Число %d, находится в масиве под индексом %d, поиск произведен за %d шага\n", num, result, count)
	}

}
