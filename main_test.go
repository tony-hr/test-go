package main

import (
	"reflect"
	"testing"
)

var (
	terbesar        Terbesar = *NewTerbesar([]int{1, 2, 3, 8, 9, 3, 2, 1})
	sortExpectedArr          = []int{1, 2, 3}
	arrSortedInLast          = []int{1, 2, 3}
	maxInArray               = 3
)

func TestUrutkanArray(t *testing.T) {
	t.Logf("Array sorted : %v", terbesar.SortArr())

	if reflect.DeepEqual(terbesar.SortArr(), sortExpectedArr) == false {
		t.Errorf("SALAH! seharusnya %v", sortExpectedArr)
	}
}

func TestArrayBerderetTerakhir(t *testing.T) {
	t.Logf("Check array sorted in the last : %v", terbesar.CheckArrSortedInLast(terbesar.SortArr()))

	if reflect.DeepEqual(terbesar.CheckArrSortedInLast(terbesar.SortArr()), arrSortedInLast) == false {
		t.Errorf("SALAH! seharusnya %v", arrSortedInLast)
	}
}

func TestAngkaMaksimal(t *testing.T) {
	t.Logf("Find max value in array : %v", terbesar.FindMaxOnArray(terbesar.CheckArrSortedInLast(terbesar.SortArr())))

	if terbesar.FindMaxOnArray(terbesar.SortArr()) != maxInArray {
		t.Errorf("SALAH! seharusnya %d", maxInArray)
	}
}
