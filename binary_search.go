package main

import (
	"fmt"
)

func main() {
	s := make([]int, 128, 128)

	for i := range s {
		s[i] = i
	}

	counter, answer := binarySearch(s, 4)
	fmt.Println(counter, answer)

	var count int
	answer2 := binarySearchRecursive(s, &count, 4)
	fmt.Println(count, answer2)

	var count2 int
	answer3 := MoreGoLikeBinarySearchRecursive(s, &count2, 4)
	fmt.Println(count2, answer3)
}

// This function will end up being the most performant as we are not making
// continuous copies of the slice and passing it along.
// All functions have the same time complexity of O(log n).
func binarySearch(list []int, item int) (int, int) {
	var (
		counter int
		low     int
	)

	high := len(list) - 1

	for low <= high {
		counter++

		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {

			return counter, guess
		} else if guess > item {
			high = mid - 1
			fmt.Println("Guess too high", low, "high>", high, "mid>", mid, "guess>", guess)
		} else {
			low = mid + 1
			fmt.Println("Guess too low>>", low, "high>", high, "mid>", mid, "guess>", guess)
		}
	}

	return 0, 0
}

// This is the function we were working on. This function will end up using
// more memory resources as we will continously create copies of the slice.
// All functions have the same time complexity of O(log n).
func binarySearchRecursive(list []int, counter *int, item int) int {
	*counter++

	mid := (len(list) - 1) / 2
	guess := list[mid]

	if guess == item {
		fmt.Println("Yaaaaayy", guess, item)
		return guess
	} else if guess > item {
		halfList := list[:mid+1]
		return binarySearchRecursive(halfList, counter, item)
	} else {
		halfList := list[mid:]
		return binarySearchRecursive(halfList, counter, item)
	}
}

// This function is just a more "GO" like version of the previous function.
// All functions have the same time complexity of O(log n).
func MoreGoLikeBinarySearchRecursive(list []int, counter *int, item int) int {
	mid := (len(list) - 1) / 2

	*counter++
	switch {
	case list[mid] == item:
		return list[mid]
	case list[mid] > item:
		return binarySearchRecursive(list[:mid+1], counter, item)
	case list[mid] < item:
		return binarySearchRecursive(list[mid:], counter, item)
	default:
		fmt.Println("could not find item")
		return -1
	}
}
