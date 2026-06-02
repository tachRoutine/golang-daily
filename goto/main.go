package main
import "fmt"

func main() {
	demo(-1)
}

func demo(n int){
	if n< 0 {
		goto Negative
	}
	fmt.Println("n is positive")
	return
	Negative:
	fmt.Println("n is negative")
}
