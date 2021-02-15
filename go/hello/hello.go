package main

import "fmt"

func main() {
    fmt.Println("Hello, world")

    var a []int
    a = make([]int, 5, 8)
    a[0] = 1
}
