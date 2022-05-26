package main

import (
	"fmt"
	"practice/bubblesort"
	"practice/quicksort"
	"practice/selectionsort"
)

func main() {

	//one way to initialize an array
	//var input [5]int = [5]int{7, 1, 9, 3, 2}

	//var input[5]int
	//input = [5]int{7, 1, 9, 3, 2}
	//can also be assigned partially, rest are default
	//input = [5]int{7, 1, 9}

	//ellipsis
	//input := [...]int{7, 1, 9, 3, 2}

	//stylistically preferred
	input := []int{7, 1, 9, 3, 2}
	sorted := quicksort.Quicksort(input, 0, len(input)-1)
	fmt.Println("quicksort:", sorted)

	input = []int{9, 2, 6, 1, 2}
	sorted = bubblesort.Bubblesort(input, 0, len(input)-1)
	fmt.Println("bubblesort:", sorted)

	input = []int{78, 1, 63, 5467, 213}
	sorted = selectionsort.Selectionsort(input)
	fmt.Println("selectionsort:", sorted)
}

//https://go.dev/tour/flowcontrol/1
//https://www.golangprograms.com/go-language/arrays.html
