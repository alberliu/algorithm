package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	slice := []int{1, 3, 5, 6, 8}
	fmt.Println(BinarySearch(slice, 3))
}

func BinarySearch(slice []int, value int) int {
	var (
		start = 0
		end   = len(slice) - 1
		mid   int
	)
	for start <= end {
		mid = (start + end) / 2
		if slice[mid] == value {
			return mid
		} else if slice[mid] < value {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
