package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int32{1, 2, 3, 4, 3, 3, 2, 1}
	fmt.Print(cutTheSticks(array))
}

func cutTheSticks(arr []int32) []int32 {
	var res []int32
	sortArray(arr)

	res = append(res, int32(len(arr)))
	x, arr := arr[0], arr[1:]

	for len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			arr[i] = arr[i] - x
		}

		arr = removeAll(arr, 0)
		res = append(res, int32(len(arr)))
		x = arr[0]
	}

	return res
}

func removeAll(arr []int32, value int32) []int32 {
	var tmp []int32
	for i := 0; i < len(arr); i++ {
		if arr[i] > value {
			tmp = append(tmp, arr[i])
		}
	}
	return tmp
}

func sortArray(array []int32) []int32 {
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	return array
}
