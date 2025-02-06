package quicksort

// QuickSort sortiert die gegebene Liste mit dem Quicksort-Algorithmus.
func QuickSort(list []int) {
	if len(list) < 2 {
		return
	}

	// HINWEIS:
	// Wählen Sie das erste Element der Liste als Pivot-Element.
	// Erstellen Sie dann zwei Listen:
	// - eine Liste mit allen Elementen, die kleiner als das Pivot-Element sind
	// - eine Liste mit allen Elementen, die größer oder gleich dem Pivot-Element sind
	// Rufen Sie QuickSort rekursiv für beide Listen auf.
	// Fügen Sie die sortierten Listen zusammen, um die sortierte Liste zu erhalten.

	pivot := list[0]
	left := []int{}
	right := []int{}

	for _, v := range list[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	QuickSort(left)
	QuickSort(right)

	copy(list, append(append(left, pivot), right...))
}
