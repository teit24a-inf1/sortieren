package bubblesort

// BubbleSort sortiert the gegebene Liste mit dem Bubble-Sort-Algorithmus.
func BubbleSort(list []int) {
	// HINWEIS:
	// Verwenden Sie eine For-Schleife, um über die Liste zu iterieren.
	// Rufen Sie BubbleUp für die Liste auf, beginnend mit der gesamten Liste
	// und dann nach und nach mit einer kürzeren Liste.
	for i := len(list) - 1; i > 0; i-- {
		BubbleUp(list[:i+1])
	}
}
