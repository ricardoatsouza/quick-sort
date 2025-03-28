package main

import (
	"fmt"
	"ricardoatsouza/go-quick-sort/src/mysort"
)

func main() {
	arr := []float64{5, 3, 7, 2, 1, 43.1, 9, -9, -5, 11, 43}

	sortedArr := mysort.QuickSort(arr)

	fmt.Println(arr)
	fmt.Println(sortedArr)

}
