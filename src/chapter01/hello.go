package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	num := []int{1, 2, 3, 4, 5, 6}
	//前2个元素
	fmt.Println(num[:2])
	//从第2个元素之后的所有元素
	fmt.Println(num[2:])
}
