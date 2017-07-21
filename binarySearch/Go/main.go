//Given a sorted array arr[] of n elements, write a function to search a given element x in arr[]

// Binary Search: Search a sorted array by repeatedly dividing the search interval in half.
// Begin with an interval covering the whole array.
// If the value of the search key is less than the item in the middle of the interval, narrow the interval to the lower half.
// Otherwise narrow it to the upper half. Repeatedly check until the value is found or the interval is empty.

//Example: [10]array{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}, searching for 23
//Start search at midpoint: 16, result 23 > 16
//Take midpoint of second half, 23 to 91
//Search new midpoint: result 23 < 56
//Two remaining, if not first number, take the second.

package main

import "fmt"

func main() {

	a := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91, 122, 432, 444, 593, 1123}
	find := 56

	fmt.Println("Running Program")
	fmt.Println("Searching list of numbers: ", a)
	fmt.Println("Searching for number: ", find)

	// numFound := false
	result, count := binarySearch(a, find)
	if result == -1 {
		fmt.Println("Your number was not found in this list.")
	} else {
		fmt.Println("Found! Your number is found at index: ", result)
		fmt.Println("Your search required ", count, " cycles with the Binary method.")
	}
}

func binarySearch(a []int, find int) (result int, count int) {
	// count = 0
	// low := a[0]
	// high := a[len(a)-1]
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 //Not found
	case a[mid] > find:
		//: is sub-slicing. means that slice from beginning up to mid
		result, count = binarySearch(a[:mid], find)
	case a[mid] < find:
		//: is sub-slicing. means from index (mid+1) upward to end of slice
		result, count = binarySearch(a[mid+1:], find)
		result += mid + 1
	default: //find == a[mid]:
		result = mid
	}
	count++
	return
}
