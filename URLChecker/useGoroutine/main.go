package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c) // "nico": true
	fmt.Println(<-c) // "flynn": true
	// fmt.Println(<-c) 없는 메시지를 받으려고 하면 에러 발생
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
