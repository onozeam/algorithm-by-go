package main

import "fmt"

// 二分探索
func main() {
	var target = 8
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var center = len(arr) / 2
	fmt.Println(arr, center)
	for true {
		if target > arr[center] {
			arr = arr[center:]
			center = len(arr) / 2
			fmt.Println(arr, center)
		} else if target < arr[center] {
			arr = arr[:center]
			center = len(arr[:center]) / 2
			fmt.Println(arr, center)
		} else {
			fmt.Println("found", arr[center])
			break
		}
	}
}
