package bubblesort

//Steps for bubblesort:
//1.  Iterate through the input array looking at two elements at a time
//2.  If the two elements are not in order, swap them
//3.  Continue this process until no swaps are made.  The array is then in order.

//Time: O(n^2)
//Space: O(1), it's in-situ
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