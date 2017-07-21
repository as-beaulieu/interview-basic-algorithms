package main

import "fmt"

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted Array:", a)
	bubbleSort(a)
	fmt.Println("Sorted Array: ", a)
}

func bubbleSort(a []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(a)-1; i++ {
			if a[i+1] < a[i] {
				swap(a, i, i+1)
				swapped = true
			}
		}
	}
}

func swap(a []int, i, j int) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}
