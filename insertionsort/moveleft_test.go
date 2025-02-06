package insertionsort

import (
	"sortieren/testhelpers"
	"testing"
)

func TestMoveLef(t *testing.T) {
	list := []int{3, 2, 1}
	MoveLeft(list, 2)
	testhelpers.AssertListsEqual(t, []int{1, 3, 2}, list)
}
