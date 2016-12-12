package selectionsort

import "fmt"

func SelectionSort(values []int)  {
	fmt.Println("This is a selection sort")
	//算法复杂度 O(n^2)

	//循环比不变量，每次循环过程中，这一结构属性都为真
	/*前缀部分，后后缀部分，
	前缀部分有序，后缀部分无序，
	循环从前缀部分没有，每次在其中加入一个元素，直到没有后缀部分，则前缀部分变得有序*/
	for i := 0; i < len(values); i++ {
		minValue := values[i]
		minIndex := i

		//从j开始后面为无序部分，选择无序部分最小值放入有序部分
		//每次循环只做一次交换
		for j := i + 1; j <len(values); j++ {
			if ( values[j] < minValue) {
				minValue = values[j]
				minIndex = j
			}
		}

		values[i], values[minIndex] = values[minIndex], values[i]
	}
}
