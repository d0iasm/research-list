package main

import (
	"fmt"
	"time"
)

var c, python, java bool

const Pi = 3.14

type Vertex struct {
	X int
	Y int
}

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

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	go say("world")
	say("hello")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	i, j := 42, 2701

	p := &i
	fmt.Println("read i through the pointer: ", *p)
	fmt.Println("see i pointer: ", p)

	*p = 21
	fmt.Println("see the new value of i: ", i)

	p = &j
	*p = *p / 37
	fmt.Println("see the new value of j: ", j)

	fmt.Println(Vertex{1, 2})

	defer fmt.Println("defer_world")
	fmt.Println("hello")

	sum := 0
	for k := 0; i < 10; i++ {
		sum += k
	}
	fmt.Println(sum)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	var n int
	k := 3
	fmt.Println(n, k, c, python, java)

	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
