package main

import (
    "fmt"
    "reflect"
)

package main

import (
    "fmt"
    "reflect"
)

func main() {
    b := true
    s := ""
    n := 1
    f := 1.0
    a := []string{"foo", "bar", "baz"}

    fmt.Println(reflect.ValueOf(b).Kind())
    fmt.Println(reflect.ValueOf(s).Kind())
    fmt.Println(reflect.ValueOf(n).Kind())
    fmt.Println(reflect.ValueOf(f).Kind())
    fmt.Println(reflect.ValueOf(a).Index(0).Kind()) // For slices and strings
}

// I found 3 ways to recognize type at runtime:

// Using string formatting

// func typeof(v interface{}) string {
//     return fmt.Sprintf("%T", v)
// }

// Using reflect package

// func typeof(v interface{}) string {
//     return reflect.TypeOf(v).String()
// }

// Using type assertions

// func typeof(v interface{}) string {
//     switch t := v.(type) {
//     case int:
//         return "int"
//     case float64:
//         return "float64"
//     //... etc
//     default:
//         _ = t
//         return "unknown"
//     }
// }
// Every method has a different best use case:

// string formatting - short and low footprint (not necessary to import reflect package)
// reflect package - when need more details about the type we have access to the full reflection capabilities
// type assertions - allows grouping types, for example recognize all int32, int64, uint32, uint64 types as "int"