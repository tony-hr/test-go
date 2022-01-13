package main

import (
	"fmt"
)

func main() {
	source := []int{1, 2, 3, 8, 9, 3, 2, 1}

	//Filter array berderet
	arrSorted := sortArr(source)

	//Cek array berderet di  array source dengan urutan sebaliknya
	result := checkArrSortedInLast(arrSorted, source)

	//Mencari nilai tertinggi dari array result
	max := result[0]
	for _, val := range result {
		if val > max {
			max = val
		}
	}

	fmt.Printf("Nilai terbesar %d (dari deret %v)", max, result)
}

func sortArr(arr []int) []int {
	var arrSorted []int

	for i, item := range arr {
		if i+1 < len(arr) {
			if (item + 1) == arr[i+1] {
				arrSorted = append(arrSorted, item)
			} else {
				if len(arrSorted) != 0 {
					arrSorted = append(arrSorted, item)
					break
				}
			}
		}
	}

	return arrSorted
}

func checkArrSortedInLast(arrSorted, source []int) []int {
	var arrFix []int
	j := 0

	for i := len(source); i > 0; i-- {
		if source[i-1] == arrSorted[j] {
			arrFix = append(arrFix, source[i-1])
		}

		j++
		if j >= len(arrSorted) {
			break
		}
	}

	return arrFix
}
