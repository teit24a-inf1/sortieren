package mergesort

// MergeSort sortiert eine Liste von Zahlen mit dem Mergesort-Algorithmus.
func MergeSort(list []int) {
	if len(list) < 2 {
		return
	}

	// HINWEIS:
	// Berechnen Sie den Index für die Mitte der Liste
	// und teilen Sie die Liste in zwei Hälften.
	// Rufen Sie MergeSort rekursiv für beide Hälften auf
	// und führen Sie dann Merge für die beiden sortierten Hälften aus.

	mid := len(list) / 2
	left := list[:mid]
	right := list[mid:]

	MergeSort(left)
	MergeSort(right)

	result := Merge(left, right)
	copy(list, result)
}
