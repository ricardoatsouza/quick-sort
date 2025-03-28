package mysort

func swap(arr []float64, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []float64, start int, end int) int {
	pivot_index := end - 1
	pivot_number := arr[pivot_index]

	i := start - 1
	for j := start; j < pivot_index; j++ {
		if arr[j] < pivot_number {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, i+1, pivot_index)
	return i + 1

}

func quick(arr []float64, start int, end int) {
	if start < end {
		pivot_index := partition(arr, start, end)

		quick(arr, start, pivot_index)
		quick(arr, pivot_index+1, end)
	}
}

func QuickSort(arr []float64) []float64 {
	arr2 := make([]float64, len(arr))
	copy(arr2, arr)

	quick(arr2, 0, len(arr))
	return arr2
}
