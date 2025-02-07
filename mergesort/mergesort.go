package mergesort

// MergeSort sortiert eine Liste von Zahlen mit dem Mergesort-Algorithmus.
func MergeSort(list []int) {
	if len(list) < 2 {
		return
	}

	m := len(list) / 2
	links := append([]int{}, list[:m]...)
	rechts := append([]int{}, list[m:]...)

	MergeSort(links)
	MergeSort(rechts)

	result := Merge(links, rechts)
	copy(list, result)

	// TODO
}
