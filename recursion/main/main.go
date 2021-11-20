package main

import "fmt"

func main()  {

	arr := make([]int, 0)
	//arr = append(arr, 11, 8, 3, 9, 7, 1, 2, 5)
	arr = append(arr, 11, 8, 3, 9)

	MergeSort(arr)

}

func MergeSort(arr []int) {
	arrLen := len(arr)
	mergeSort(arr, 0, arrLen - 1)

}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	fmt.Println(start, end, mid)

	mergeSort(arr, start, mid)

	fmt.Println("aaaaa")
	fmt.Println(start, end, mid)
	fmt.Println("bbbbb")

	mergeSort(arr, mid + 1, end)


}