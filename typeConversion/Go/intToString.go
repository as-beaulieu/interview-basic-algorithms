// You can concat strings simply by +'ing them, or by using the Join function of the strings package.

package main

import (
    "strconv"
    "fmt"
)

func main() {
    t := strconv.Itoa(123)
    fmt.Println(t)
}

// It is interesting to note that strconv.Itoa is shorthand for

// func FormatInt(i int64, base int) string
// with base 10

// For Example:

// strconv.Itoa(123)
// is equivalent to
// strconv.FormatInt(int64(123), 10)