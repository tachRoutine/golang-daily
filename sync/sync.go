package main

import "sync"

var count = 100
var wg sync.WaitGroup

func printA(){
	for range count{
		print("-A-")
	}
	wg.Done()
}

func printB(){
	for range count{
		print("-B-")
	}
	wg.Done()
}

func printC(){
	for range count{
		print("-C-")
	}
	wg.Done()
}

func main(){
	defer wg.Wait()
	wg.Add(3)
	go printA()
	go printB()
	go printC()
}
