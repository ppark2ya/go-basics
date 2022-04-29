package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [4]string{"nico", "flynn", "jtpark", "larry"}

	for _, person := range people {
		go isSexy(person, channel)
	}

	for i := 0; i < len(people); i++ {
		// blocking operation -> 결과를 받기 전까지 for문이 멈춰있는다.
		fmt.Println(<-channel)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}
