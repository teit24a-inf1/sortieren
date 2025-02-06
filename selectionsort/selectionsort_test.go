package selectionsort

import (
	"sortieren/testhelpers"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := []int{3, 2, 1}
	SelectionSort(list)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3}, list)
}
