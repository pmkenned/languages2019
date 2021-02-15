package main

import (
	"fmt"
	"time"
	"math"
	"errors"
)

type person struct {
	name string
	age int
}

func main() {
	var x int
	x=5
	y := 2
	fmt.Println(x+y)

	var a [5]int

	b := [5]int{5,4,3,2,1}

	c := []int{7,8,2,6,1}

	fmt.Println(a)
	fmt.Println(b)

	if a[0] == 5 {
		fmt.Println("a[0] == 5")
	} else if a[0] == 0 {
		fmt.Println("a[0] ==0")
	} else {
		fmt.Println("a[0] != 0 or 5")
	}

	c = append(c, 13)
	fmt.Println(c)

	fmt.Println("The time is", time.Now())

	vertices := make(map[string]int )
	vertices["traingle"] = 3
	vertices["square"] = 4
	fmt.Println(vertices)
	delete(vertices, "square")
	fmt.Println(vertices)

	for i:= 0; i<5; i++ {
		fmt.Println(i)
	}

	i:= 0
	for i<5 {
		fmt.Println(i)
		i++
	}

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	result := sum(2,3)
	fmt.Println(result)

	result2, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	p := person{name: "Jake", age: 25}
	fmt.Println(p.age)
	z := 7
	fmt.Println(&z)
	fmt.Println(z)
	inc(&z)
	fmt.Println(z)
}

func sum (x int, y int) int {
	return x+y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func inc(x *int) {
	*x++
}
