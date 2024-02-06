package main

import (
	"errors"
	"fmt"
	"sort"
)

func sortSlice(ar []int) []int {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	return ar
}

func binarySearch(min, max, searchItem int, sortedSlice []int) (int, error) {
	if min > max {
		return 0, errors.New("элемент не найден")
	}

	half := (max-min)/2 + min

	switch {
	case sortedSlice[half] == searchItem:
		return half, nil
	case sortedSlice[half] > searchItem:
		return binarySearch(min, half-1, searchItem, sortedSlice)
	default:
		return binarySearch(half+1, max, searchItem, sortedSlice)
	}
}

func main() {
	a := []int{1, 2, 5, 7, 3, 4, 6, 9, 8}
	searchIndex := 3
	s := sortSlice(a)
	index, err := binarySearch(0, len(s)-1, searchIndex, s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(index)
	}
}
