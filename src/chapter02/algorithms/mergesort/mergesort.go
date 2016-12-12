package mergesort

import "fmt"

func merge(left []int, right []int) []int {
	result := make([]int, len(left) + len(right));
	//fmt.Println(result)
	index := 0
	i, j := 0, 0
	for {
		//等长的部分，比较大小后插入
		if i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				result[index] = left[i]
				index++
				i++
			} else {
				result[index] = right[j]
				index++
				j++
			}
		} else {
			break
		}
	}

	//以下处理元素个数为非2^2的情况，左边或者右边多一个但一定是最大的直接追加在尾部
	for {
		//左边比右边长
		if i < len(left) {
			result[index] = left[i]
			index++
			i++
		} else {
			break
		}
	}

	for {
		//右边比左边长
		if j < len(right) {
			result[index] = right[j]
			index++
			j++
		} else {
			break
		}
	}
	return result[:]
}

func mergeSort(values []int) []int{
	//fmt.Println(values)
	if len(values) < 2 {
		return values[:]
	} else {
		middle := len(values) / 2
		//数组切片，前middle个元素
		left := mergeSort(values[:middle])
		//数组切片，第middle个元素之后的所有元素
		right := mergeSort(values[middle:])
		together := merge(left, right)
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println(together)
		return together[:]
	}
	return nil
}

func MergeSort(values []int) {
	//算法复杂度 O(n·logn) 分治算法之一，每层都是n个元素，总共logn层，合计复杂度O(n·logn)
	/*
	                     n
	                 /     \
	               /        \
	            n/2           n/2
	           /   \            /  \
	         /      \          /    \
	      n/4       n/4     n/4     n/4
	      /\        /\      /  \      / \
	     /  \      /  \    /   \    /   \
	  node node node node node node node node

	*/

	fmt.Println("This is merge sort")
	result := mergeSort(values)
	for i := 0; i<len(result); i++ {
		values[i] = result[i]
	}
}
