package main

import "fmt"

// import "time"

func channeltest() {
	g := make(chan int)
	quit := make(chan bool)

	// 	go B()

	go func() {
		for {
			select {
			case i := <-g:
				// fmt.Println("go func", <-g)
				fmt.Println("go func", i+1)
			case <-quit:
				fmt.Println("B quit")
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		g <- i
	}
	quit <- true // 沒辦法等待B的退出只能Sleep
	// 	time.Sleep(time.Second * 1)
	fmt.Println("Main quit")
}

// func B() {
// 	for {
// 		select {
// 		case i := <-g:
// 			fmt.Println(i + 1)
// 		case <-quit:
// 			fmt.Println("B quit")
// 			return
// 		}
// 	}
// }
