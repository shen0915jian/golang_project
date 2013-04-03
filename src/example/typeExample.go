// typeExample
package main

import (
	"fmt"
)

type testInt func(int) bool //声明一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

//声明的函数类型当作一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice= ", slice)

	odd := filter(slice, isOdd) //函数当值来传递
	fmt.Println("Odd element of slice are: ", odd)

	even := filter(slice, isEven) //函数当值来传递
	fmt.Println("Even element of slice are: ", even)
}
