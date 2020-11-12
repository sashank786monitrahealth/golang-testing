package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	var elements []int = []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	BubbleSort(elements)

}
