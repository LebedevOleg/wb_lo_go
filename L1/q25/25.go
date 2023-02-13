package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	var stop time.Duration = 5 * time.Second
	fmt.Println("время ", time.Now())
	sleep(stop)
	fmt.Println("время ", time.Now())

}
