package myslice

import "fmt"

func MyArray() {
	arr1 := []int{1, 2}
	arr2 := arr1[:]
	arr3 := make([]int, 5)
	checkAddr(&arr1)
	checkAddr(&arr2)
	fmt.Printf("arr1 : %p, %v\n", &arr1, arr1)
	fmt.Printf("arr2 : %p, %v\n", &arr2, arr2)
}

func checkAddr(arr *[]int) {
	fmt.Printf("func Array : %p , %v\n", &arr, arr)
}