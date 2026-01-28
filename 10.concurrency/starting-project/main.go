package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
}

func main() {
	done := make(chan bool)

	greet("Nice to meet you!", done)
	greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	greet("I hope you're liking the course!", done)

	for range done {
		fmt.Println("done")
	}
}
