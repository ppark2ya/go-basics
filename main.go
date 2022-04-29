package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutines -> main 함수가 동작하는 동안 병렬로 실행
	go sexyCount("nico")
	go sexyCount("jtpark")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
