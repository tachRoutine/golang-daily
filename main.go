package main

import (
	"fmt"
	"time"
)

var Weekday string

func init(){
	Weekday = time.Now().Weekday().String()
}

func main(){
	fmt.Println("Today is",Weekday)
}
