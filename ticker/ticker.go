package ticker

import (
	"fmt"
	"time"
)

func init(){
	fmt.Println("This is the ticker pkg")
}

func Tickeeerrr(){
	ticker := time.NewTicker(5*time.Second)
	
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
	
}