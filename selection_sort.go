package main

import (
	"fmt"
	"math/rand"
)

func main() {
	unsorted := make([]int, 5)
	for i := range unsorted {
		unsorted[i] = rand.Intn(1000)
	}
	fmt.Println("Original:", unsorted)

	sorted := make([]int, 5)

	_, _, finalSorted := findSmallest(0, unsorted, sorted)
	fmt.Println("Sorted Final:", finalSorted)
}

func findSmallest(index int, unsorted, sorted []int) (int, []int, []int) {
	if len(unsorted) == 0 {
		return index, unsorted, sorted
	}

	var smallestIndex int
	smallest := unsorted[len(unsorted)-1]

	for i, v := range unsorted {
		if v <= smallest {
			smallestIndex, smallest = i, v
		}
	}

	sorted[index] = smallest

	newUnsorted := make([]int, 0)
	newUnsorted = append(newUnsorted, unsorted[:smallestIndex]...)
	newUnsorted = append(newUnsorted, unsorted[smallestIndex+1:]...)

	index++

	return findSmallest(index, newUnsorted, sorted)
}
