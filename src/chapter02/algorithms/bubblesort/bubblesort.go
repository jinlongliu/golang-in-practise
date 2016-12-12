package bubblesort

import "fmt"

func BubbleSort(values []int) {
	fmt.Println("This is a bubble sort")
	//算法复杂度O(n^2)  n = len(values)， 通选择排序相比复杂度相同，但是选择排序可能更加高效，每次循环中交换次数不同
	flag := true

	for i := 0; i<len(values) -1; i++ {
		flag = true

		//该循环每执行一次确保最大的沉底，大的往后冒泡
		//每次循环中多次交换
		for j := 0; j < len(values) - i - 1; j++ {
			if values[j] > values[j + 1] {
				values[j], values[j + 1] = values[j + 1], values[j]
				flag = false
			}//end if
		}//end for j = ...

		//在一次循环中如果没有任何位置调整，说明所有数值前者大于后者，排序提前结束
		if flag == true {
			break
		}
		//fmt.Println(i)
	}//end for i = ...
}
