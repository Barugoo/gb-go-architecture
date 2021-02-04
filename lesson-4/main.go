package main

import (
	"flag"
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func Search(a []int, search int) (result int, count float64) {
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
	var counts []float64

	var arr [][]int

	arr[0] = []int{10, 20, 30, 40, 50}

	searchArray0 := []int{10, 20, 30, 40, 50}
	searchArray1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	searchArray2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
	searchArray3 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}
	searchArray4 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250}
	searchArray5 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300}
	flag.IntVar(&num, "n", 0, "Search number")

	flag.Parse()

	fmt.Printf("Поиск числа %d\n", num)

	_, count := Search(searchArray0, num)
	counts = append(counts, count)

	_, count = Search(searchArray1, num)
	counts = append(counts, count)

	_, count = Search(searchArray2, num)
	counts = append(counts, count)

	_, count = Search(searchArray3, num)
	counts = append(counts, count)

	_, count = Search(searchArray4, num)
	counts = append(counts, count)

	_, count = Search(searchArray5, num)
	counts = append(counts, count)

	graph := asciigraph.Plot(counts)

	fmt.Println(graph)
}
