package bubblesort

func Bubblesort(array []int, start, end int) []int{
	
	var swap bool
	swap = true

	for ; swap ; {
		swap = false

		for i := 1; i < len(array); i++ {
			if array[i-1] > array[i] {
				swap = true
				array[i-1], array[i] = array[i], array[i-1]
			}
		}
	}

	return array
}