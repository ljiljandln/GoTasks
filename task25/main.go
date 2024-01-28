package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	// горутина блокируется на время равное d
	<-time.After(d)
}

func main() {
	start := time.Now()
	mySleep(5 * time.Second)
	fmt.Println(time.Since(start))
}
