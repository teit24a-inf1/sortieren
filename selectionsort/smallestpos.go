package selectionsort

// SmallestPos gibt die Position des kleinsten Elements in der Liste zurück.
func SmallestPos(list []int) int {
	smallest := 0

	// HINWEIS:
	// Verwenden Sie eine For-Schleife, um über die Liste zu iterieren.
	// Vergleichen Sie das aktuelle Element mit dem kleinsten Element.
	// Aktualisieren Sie die Position des kleinsten Elements, wenn das aktuelle Element kleiner ist.

	for i := 1; i < len(list); i++ {
		if list[i] < list[smallest] {
			smallest = i
		}
	}
	return smallest
}
