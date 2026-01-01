package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	signals <- true

	select {
	case msg := <-messages:
		fmt.Println("recived message", msg)
	default:
		fmt.Println("no message recived")
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("recived message", msg)
	case sig := <-signals:
		fmt.Println("recived signal", sig)
	default:
		fmt.Println("no activity")
	}
}
