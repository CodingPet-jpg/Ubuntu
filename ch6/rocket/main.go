package main

import (
	"fmt"
	"time"
)

type Rocket struct {
	Z int
	S string
}

func (rocket Rocket) Launch(i int) {
	fmt.Printf("Launch After %d s\n", i)
}

func main() {
	r := new(Rocket)
	i := 3
	timer := time.AfterFunc(time.Second*time.Duration(i), func() {
		r.Launch(i)
	})
	fmt.Println("Can i land?")
	time.Sleep(5 * time.Second)
	i = 6 // no effect
	var log string = "Success"
	if timer.Stop() {
		log = "Failed"
	}
	fmt.Println(log)
	fmt.Println("Hello Moon")
}
