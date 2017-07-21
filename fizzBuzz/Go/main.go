//Fizz Buzz

//Write a program that prints the numbers from 1 to 100.
//But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
//For numbers which are multiples of both three and five print “FizzBuzz”.

//Result:
// 1
// 2
// fizz
// 4
// buzz
// fizz
// 7
// 8
// fizz
// buzz
// 11
// fizz
// 13
// 14
// fizz buzz
// 16
// 17
// fizz
// 19
// buzz
// etc... to 100

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}

//Another method:
func moduloFifteen() {
	for i := 1; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Println("Fizz-Buzz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else if i%3 == 0 {
			fmt.Println("fizz", i)
		} else {
			fmt.Println(i)
		}
	}
}

//Using a Switch instead of if else:
func switchMethod() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
