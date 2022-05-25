package main

import (
	"fmt"
	"practice/quicksort"
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

	fmt.Println("Pre-sorted array")
	fmt.Println(input)

	fmt.Println("quicksort")
	sorted := quicksort.Quicksort(input, 0, len(input) - 1)
	fmt.Println(sorted)
}

//https://go.dev/tour/flowcontrol/1
//https://www.golangprograms.com/go-language/arrays.html