package main

import (
	"fmt"
	"practice/ticker"
	"time"
)

var Weekday string

func init(){
	Weekday = time.Now().Weekday().String()
	ticker.Tickeeerrr()
}

func main(){
	fmt.Println("Today is",Weekday)
}
