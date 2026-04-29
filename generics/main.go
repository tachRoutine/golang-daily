package main

import "fmt"

func main() {
	print(100)
	print("Hello, Generics!")
}

func print[T any](value T) {
	fmt.Println(value)
}
