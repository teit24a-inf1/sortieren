package mergesort

// Merge erwartet zwei sortierte Listen und gibt eine sortierte Liste zurück,
// die alle Elemente der beiden Eingabelisten enthält.
func Merge(list1, list2 []int) []int {
	result := []int{}

	// HINWEIS:
	// Verwenden Sie eine For-Schleife, um über die beiden Listen zu iterieren.
	// Vergleichen Sie die beiden jeweils ersten Elemente der Listen.
	// Fügen Sie das kleinere Element zur Ergebnisliste hinzu
	// und entfernen Sie es aus der ursprünglichen Liste.
	for len(list1) > 0 && len(list2) > 0 {
		if list1[0] < list2[0] {
			result = append(result, list1[0])
			list1 = list1[1:]
		} else {
			result = append(result, list2[0])
			list2 = list2[1:]
		}
	}
	result = append(result, list1...)
	result = append(result, list2...)

	return result
}
