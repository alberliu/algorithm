package algorithm

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	slice := []int{5, 4, 7, 2, 6, 3, 6, 9, 19}
	//slice := []int{4, 5, 2, 6, 3}

	QuickSort(slice)
	fmt.Println(slice)
}

// QuickSort 快速排序
func QuickSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	index := partition(slice)
	QuickSort(slice[:index])
	QuickSort(slice[index+1:])

}

func partition(slice []int) int {
	index := 0
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			index = i + 1
		}
	}
	return index
}
