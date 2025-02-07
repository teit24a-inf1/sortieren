package bubblesort

// BubbleUp implementiert eine Iteration des Bubble-Sort-Algorithmus.
// Es iteriert über die Liste und vergleicht jedes Element mit dem nächsten.
// Wenn das Element größer als das nächste ist, werden sie vertauscht.
// Die Funktion gibt true zurück, wenn mindestens ein Tausch durchgeführt wurde.
func BubbleUp(list []int) bool {
	swapped := false

	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
			swapped = true
		}
	}

	return swapped
}
