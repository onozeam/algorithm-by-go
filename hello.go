package main

import "fmt"

// 二分探索
func main() {
	var target = 2
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var center = len(arr) / 2
	fmt.Println(arr[center])
	for true {
		if target > arr[center] {
			fmt.Println(">")
		} else if target < arr[center] {
			fmt.Println("<")
		} else {
			fmt.Println("o")
		}
		break
	}
}
