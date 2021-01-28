package main

import (
	"flag"
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

	var num int

	searchArray0 := []int{10, 20, 30, 40, 50}
	searchArray1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	searchArray2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
	searchArray3 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}

	flag.IntVar(&num, "n", 0, "Search number")

	flag.Parse()

	fmt.Printf("Поиск числа %d\n", num)

	result, count := Search(searchArray0, num)
	fmt.Printf("Число %d, находится в масиве searchArray0 под индексом %d, поиск произведен за %d шага\n", num, result, count)
	result, count = Search(searchArray1, num)
	fmt.Printf("Число %d, находится в масиве searchArray1 под индексом %d, поиск произведен за %d шага\n", num, result, count)
	result, count = Search(searchArray2, num)
	fmt.Printf("Число %d, находится в масиве searchArray2 под индексом %d, поиск произведен за %d шага\n", num, result, count)
	result, count = Search(searchArray3, num)
	fmt.Printf("Число %d, находится в масиве searchArray3 под индексом %d, поиск произведен за %d шага\n", num, result, count)
}
