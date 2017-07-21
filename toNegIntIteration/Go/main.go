//2 non-negative integers

//There are 2 non-negative integers: i and j.
//Given the following equation, find an (optimal) solution to iterate over i and j in such a way that the output is sorted.
//2^i * 5^j

// So the first few rounds would look like this:
// 2^0 * 5^0 = 1
// 2^1 * 5^0 = 2
// 2^2 * 5^0 = 4
// 2^0 * 5^1 = 5
// 2^3 * 5^0 = 8
// 2^1 * 5^1 = 10
// 2^4 * 5^0 = 16
// 2^2 * 5^1 = 20
// 2^0 * 5^2 = 25

//Dijkstra derives an eloquent solution in "A Discipline of Programming".
//He attributes the problem to Hamming. Here is my implementation of Dijkstra’s solution.

package main

import "fmt"

func main() {
	//Generate the first n numbers
	n := 20
	var v []int
	//v[0] = 1

	//Index for 2
	i2 := 0
	//Index for 5
	i5 := 0

	//Next two candidates
	x2 := 2 * v[i2]
	x5 := 5 * v[i5]

	for i := 1; i >= n; i++ {
		m, err := fmt.Println(x2, x5)
		if err != nil {
			fmt.Println("Error", err)
		}
		v[i] = m

		if x2 == m {
			i2++
			x2 = 2 * v[i2]
		}
		if x5 == m {
			i5++
			x5 = 5 * v[i5]
		}
	}

}


//PROBLEM: CODE DOESNT WORK: PANIC AT LINE 29!!!!