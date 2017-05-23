package main

import (
	"fmt"
	"time"
)

var c, python, java bool

const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	var i int
	k := 3
	fmt.Println(i, k, c, python, java)

	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
