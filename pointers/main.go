package main

import "fmt"

// Creating pointers
// Pointers are used to modify data that exists outside of a function
// Asterisk (*) when used with a type indicates the value is a pointer
// Ampersand (&) creates a pointer from a variable

// Asterisk (*) on a variable will dereference the pointer

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

	fmt.Println(hello, world)
}
