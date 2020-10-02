package main

import (
	"fmt"
	"time"
)

func main() {

	// channel
	c := make(chan string)

	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}

	// wait for channel, blocking operation
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for", i)
		fmt.Println(<-c)
	}
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy"
}
