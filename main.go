package main

import (
	"fmt"
)

type Terbesar struct {
	Data []int
}

func NewTerbesar(d []int) *Terbesar {
	return &Terbesar{d}
}

func main() {
	source := []int{1, 2, 3, 8, 9, 3, 2, 1}
	p := NewTerbesar(source)

	//Filter array berderet
	arrSorted := p.SortArr()

	//Cek array berderet di  array source dengan urutan sebaliknya
	result := p.CheckArrSortedInLast(arrSorted)

	//Mencari nilai tertinggi dari array result
	max := p.FindMaxOnArray(result)

	fmt.Printf("Nilai terbesar %d (dari deret %v)", max, result)
}

func (p *Terbesar) SortArr() []int {
	var arrSorted []int

	for i, item := range p.Data {
		if i+1 < len(p.Data) {
			if (item + 1) == p.Data[i+1] {
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

func (p *Terbesar) CheckArrSortedInLast(arrSorted []int) []int {
	var arrFix []int
	j := 0

	for i := len(p.Data); i > 0; i-- {
		if p.Data[i-1] == arrSorted[j] {
			arrFix = append(arrFix, p.Data[i-1])
		}

		j++
		if j >= len(arrSorted) {
			break
		}
	}

	return arrFix
}

func (p *Terbesar) FindMaxOnArray(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}
