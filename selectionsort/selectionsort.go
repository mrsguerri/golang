package selectionsort

//1.  Let the array be divided into two parts, a sorted an unsorted one.
//2.  The sorted part 's' is at first denoted by position 0 of the array and length 0.
//3.  The unsorted part is the remainder of the array.
//            [()(5 8 12 8 47 0)]
//2.  Find the smallest number in the unsorted array and move it to the end of the sorted array
//3.  Repeat until the sorted array is the length of the entire array.

//Returns the minimal
func min(array []int) int {
	m := 0
	for i, v := range array {
		if v < array[m] {
			m = i
		}
	}

	return m
}

func Selectionsort(array []int) []int {
	lastSorted := 0

	for i := 0; i < len(array)-1; i++ {
		minimal := min(array[lastSorted:]) + lastSorted //add last sorted because the slice's size is reduced in the min function
		array[lastSorted], array[minimal] = array[minimal], array[lastSorted]
		lastSorted++
	}

	return array
}
