package main

import "fmt"

func main() {
	fmt.Println("llll")

	val := 3.4
	arr := []float64{1.2, 2.0, 3.3, 4.4, 5.5, 6.6, 7.8}
	flag := BinarySearch(val, arr)
	fmt.Println(flag)
}

func BinarySearch(val float64, arr []float64) int {
	lenArr := len(arr)
	left, right := 0, lenArr-1
	if lenArr < 1 {
		return -1
	}
	if lenArr == 1 {
		if val >= arr[0] {
			return 1
		} else {
			return -1
		}
	}
	flag := 0
	for left < right {
		mid := left + (right-1)/2

		if val < arr[mid] {
			//如果在左边
			right = mid - 1
			if arr[left] <= val && val < arr[left+1] {
				flag = left + 1
			}
		} else if val > arr[mid] {
			left = mid + 1
			if arr[left] <= val && val < arr[left+1] {
				flag = left + 1
			}
		} else {
			//相等就是mid + 1
			return mid + 1
		}
	}
	return flag
}
