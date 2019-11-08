package main

import "fmt"

func channel() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	go func() { messages <- "ping2" }()
	msg := <-messages
	fmt.Println(msg)

	// 	chanmsg := make(chan string)
	// 	chanmsg <- "ping 2"
	// 	msg2 := <-chanmsg
	// 	fmt.Println(msg2)
}
