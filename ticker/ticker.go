package ticker

import (
	"fmt"
	"time"
)

func Tickeeerrr(){
	ticker := time.NewTicker(5*time.Second)
	
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
	
}