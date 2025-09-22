package main

import "fmt"


func main(){
	fmt.Println("Start of main function")
	printSomething("Hello, World!")
	defer printSomething("This will be printed last")
	printSomething("Goodbye, World!")
}

func printSomething(something string){
	fmt.Println(something)
}